# stay-watch-backend


envファイルをリポジトリのrootディレクトリに置く

.envファイル
```
GOLANG_PORT=8082
MYSQL_PORT=33066
GOLANG_CONTAINER_NAME=vol_golang
MYSQL_CONTAINER_NAME=vol_mysql
```

.firebase.jsonもgo/app/credentialsに置く
slackのprj_staywatchを参照

実行方法(ローカル)<br>
/stay-watch-backend/app ディレクトリに移動して下のコマンド
```
go run main.go
```


ネットワーク作成
```
docker network create vol_network
```


go+mysqlのコンテナ起動
```
make dev
```

mysqlコンテナの入り方
```
make vol_mysql
```

mysql ログイン
```
mysql -uroot -proot
```


データベース名を指定
```
use app;
```

データベースの初期化方法(ローカル)<br>
1. 現在のmysqlコンテナを削除する。（Docker Desktop だとゴミ箱マーク）
2. コンテナを作成する。
    ```
    docker-compose up
    ```
    実行するとテーブルも何もないmysqlコンテナが作られる
3. /stay-watch-backend/app ディレクトリに移動してmain.goを実行する
    ```
    go run main.go
    ```
    すると先程のmysqlコンテナに init.sql に書いてある内容のテーブル、カラム、値が入る。











