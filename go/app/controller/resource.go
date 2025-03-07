package controller

import (

	// "strconv"

	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func BackUpDB(c *gin.Context) {
	bucketName  := "staywatch-backend"	// GCSバケット名
	credentialsFile := "./credentials/stay-watch-a616f-d11c68af2c21.json"	// サービスアカウント鍵ファイルのパス

	now := time.Now()
	formatedNow := now.Format("2006-01-02_15-04-05")
	localFileName := "./mydump_sql.bk"
	gcsFileName := "backup/db_backup_" + formatedNow + ".bk"	// GCSバケットのアップロード先のパス

	// ====== DBバックアップ実行 ======
	// 実行したいシェルスクリプトのパス
	scriptPath := "./backup.sh"

	// コマンドを準備
	cmd := exec.Command("sh", scriptPath, localFileName)

	// コマンドの標準出力と標準エラー出力を取得
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// コマンドを実行
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command execution failed: %v\n%s", err, stderr.String())
		c.JSON(http.StatusServiceUnavailable, "featal upload file to CloudStorage")
		return
	}

	// == Cloud Storageへ保存
	// Google Cloud Storageクライアントの作成
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusServiceUnavailable, "featal create CloudStorage client")
		return
	}

	// アップロードするファイルを開く
	file, err := os.Open(localFileName)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusServiceUnavailable, "featal open local file")
		return
	}
	defer file.Close()

	// バケットオブジェクトの作成
	bucket := client.Bucket(bucketName)

	// バケット内のアップロード先のオブジェクトを作成
	obj := bucket.Object(gcsFileName)

	// ファイルをアップロード
	wc := obj.NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusServiceUnavailable, "featal upload file to CloudStorage")
		return
	}
	if err := wc.Close(); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusServiceUnavailable, "featal upload file to CloudStorage")
		return
	}	

	c.JSON(http.StatusOK, "sucess")
}