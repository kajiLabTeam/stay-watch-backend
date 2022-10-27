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












