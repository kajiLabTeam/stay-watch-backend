# 環境変数
CONTAINER_NAME="vol_mysql"        # Dockerコンテナ名
CONTAINER_USER="root"
DB_USER="root"                    # MySQLユーザー名
DB_PASSWORD="root"       # MySQLパスワード
DB_NAME="app"                     # バックアップ対象のデータベース名
DUMP_PATH_IN_CONTAINER="/tmp/mydump_sql.dump"  # コンテナ内の一時保存先
DUMP_PATH_ON_HOST=$1          # ホストに保存する場所

# mysqldumpをコンテナ内で実行
ssh -i /root/.ssh/docker_rsa -o StrictHostKeyChecking=no $CONTAINER_USER@$CONTAINER_NAME "mysqldump --single-transaction -u$DB_USER -p$DB_PASSWORD $DB_NAME > $DUMP_PATH_IN_CONTAINER"
scp -i /root/.ssh/docker_rsa -o StrictHostKeyChecking=no $CONTAINER_USER@$CONTAINER_NAME:/$DUMP_PATH_IN_CONTAINER $DUMP_PATH_ON_HOST

# 結果を出力
if [ $? -eq 0 ]; then
    echo "Backup successful! Saved to $DUMP_PATH_ON_HOST"
else
    echo "Backup failed!"
fi