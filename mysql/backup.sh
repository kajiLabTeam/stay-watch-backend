# 環境変数
CONTAINER_NAME="vol_mysql"        # Dockerコンテナ名
DB_USER="root"                    # MySQLユーザー名
DB_PASSWORD="root"       # MySQLパスワード
DB_NAME="app"                     # バックアップ対象のデータベース名
DUMP_PATH_IN_CONTAINER="/tmp/mydump_sql4.dump"  # コンテナ内の一時保存先
DUMP_PATH_ON_HOST="./mysql/backup/mydump_sql4.dump"          # ホストに保存する場所

# mysqldumpをコンテナ内で実行
docker exec "$CONTAINER_NAME" bash -c "mysqldump --single-transaction -u$DB_USER -p$DB_PASSWORD $DB_NAME > $DUMP_PATH_IN_CONTAINER"

# コンテナ内からホストにダンプファイルをコピー
docker cp "$CONTAINER_NAME:$DUMP_PATH_IN_CONTAINER" "$DUMP_PATH_ON_HOST"

# 結果を出力
if [ $? -eq 0 ]; then
    echo "Backup successful! Saved to $DUMP_PATH_ON_HOST"
else
    echo "Backup failed!"
fi