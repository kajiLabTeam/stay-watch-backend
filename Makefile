include .env

# コンテナ名 envファイルから取得
DB_CONTAINER_NAME = $(MYSQL_CONTAINER_NAME)
# データベース名
DB_NAME = app
MYSQL_PASSWORD=root

# データベースをダンプ
dump:
	docker exec -it $(DB_CONTAINER_NAME) mysqldump -uroot --password=${MYSQL_PASSWORD} $(DB_NAME) > mysql/backup/tmp_backup.sql 
replace_dump: dump
	sed '1d' mysql/backup/tmp_backup.sql > mysql/backup/backup.sql
# データベースをリストア
restore:
	docker exec -i $(DB_CONTAINER_NAME) mysql -uroot -p$(MYSQL_PASSWORD) $(DB_NAME) < mysql/backup/backup.sql
	
reloadgolang:
	docker-compose rm -fsv vol_golang
	docker-compose up -d vol_golang
down:
	docker-compose down

## mysqlコンテナの立ち上げ
vol_mysql:
	docker-compose up vol_mysql

vol_mysql-d:
	docker-compose up -d vol_mysql

ex_vol_mysql:
	docker-compose exec vol_mysql bash

network:
	docker network create vol_network
# 開発環境
dev:
	docker-compose up 
dev-d:
	docker-compose up -d
# 本番環境
prod:
	docker-compose up -d vol_mysql
	sleep 120
	docker-compose up -d vol_golang
	
pull-backup-sql:
	scp  ssh-server.kajilab.tk:/home/kjlb/stay-watch-backend/mysql/backup/backup.sql ~/
