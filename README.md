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
 <br>

### データベースの初期化方法(ローカル)<br>
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

<br>

### Database の中身を閲覧、編集する方法（DBeaver）
（事前にvol_mysqlコンテナを起動しておく）

1. DBeaver をインストールし起動

2. Create sample database は cancel

3. MySQL を選択

4. 以下のように設定する
    - Server Host: localhost
    - Port: 33066
    - Database: app
    - ユーザー名: root
    - パスワード: (先駆者まで)

    他はそのまま

5. テスト接続が通ったら終了

6. 「app -> データベース -> app -> テーブル」　から閲覧、編集ができる


<br>

### 環境構築の手順 <br><br>
1. gitをインストール<br>
    homebrewが必要（「brew -v」 が使えるか）
    
    <br>
2. GitHub Desktopを入れる<br>
    <br>

3. stay-watch-backendを任意のディレクトリにクローンする<br>
    <br>

4. VSCodeでgitでコミットやプルを使えるようにしておく（任意）<br>
    VSCodeの左バーの「ソース管理」から画面に従っていけばできるはず
    
    <br>
5. .envファイルの作成<br>
    /stay-watch-backend/ のディレクトリに .envファイルを作成し、中身を先駆者からもらう

    <br>
6. firebase.json の作成<br>

    firebase.jsonを先駆者からもらう<br>

    /stay-watch-backend/go/app/ のディレクトリに credentials ディレクトリを作成し、その中にfirebase.jsonを置く
    
    <br>
7. Dockerコンテナの作成<br>
    （コマンドは全て/stay-watch-backend で行う）<br>

    DockerDesktopをインストールし起動<br>

    ネットワークを作成
    ```
    docker network create vol_network
    ```
    コンテナを作成
    ```
    docker-compose up
    ```

    <br>
8. Go のインストール<br>
    
    <br>
9. 実行<br>

    Dockerのコンテナを起動

    /Users/togawa/GitHub/stay-watch-backend/go/ のディレクトリで以下を行なってうまく動作すればOK
    ```
    go run main.go
    ```

<br>

### 必要なVSCodeの拡張機能<br><br>

- REST Client

    test.httpでAPIをテストするときに使用

- Go

    Goを扱う上であるとよい

- OpenAPI(Swagger) Editor

    swagger.yml を編集する上であるとよい

- Swagger Viewer

    Shift + command + p でswagger.yml のプレビューができる













