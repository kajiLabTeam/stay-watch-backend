include .env

# コンテナ名
CONTAINER_NAME = vol_mysql
# データベース名
DB_NAME = app
# mysqlのログイン情報を設定
set:
	docker exec -it $(CONTAINER_NAME) mysql_config_editor set -u root -p
# mysqldumpのログイン情報を設定
setdump:
	docker exec -it $(CONTAINER_NAME) mysql_config_editor set --login-path=mysqldump -u root -p
# データベースをダンプ
dump:
	docker exec -it $(CONTAINER_NAME) mysqldump $(DB_NAME) > backup.sql
# データベースをリストア
restore:
	docker exec -i $(CONTAINER_NAME) mysql $(DB_NAME) < mysql/backup/backup.sql

reloadgolang:
	docker-compose rm -fsv vol_golang
	docker-compose up -d vol_golang

down:
	docker-compose down

vol_mysql:
	docker-compose exec vol_mysql bash

dev:
	docker-compose up 
	
dev -d:
	docker-compose up -d

network:
	docker network create vol_network




	







