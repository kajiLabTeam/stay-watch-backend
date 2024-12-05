# 環境設定
CONTAINER_NAME="vol_mysql"  # Dockerコンテナ名
DB_NAME="app"              # リストア先のデータベース名
DUMP_FILE_PATH="./mysql/backup/mydump_sql4.dump"  # ダンプファイルのパス（ホスト側のファイル）

# ダンプファイルをコンテナ内にコピー
docker cp "$DUMP_FILE_PATH" "$CONTAINER_NAME:/tmp/mydump_sql1.dump"

# コンテナ内でリストアコマンドを実行
docker exec -i "$CONTAINER_NAME" bash -c "mysql -u root -p${MYSQL_ROOT_PASSWORD} $DB_NAME < /tmp/mydump_sql1.dump"

# リストア後、不要なダンプファイルを削除（任意）
docker exec -i "$CONTAINER_NAME" bash -c "rm /tmp/mydump_sql1.dump"

echo "リストアが完了しました！"