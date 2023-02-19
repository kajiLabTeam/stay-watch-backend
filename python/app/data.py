import mysql.connector


# ユーザの名前とそのユーザのログ時間の合計を返す関数
def get_log_times():
    cnx = mysql.connector.connect(
        host="vol_mysql",
        port=3306,
        database="app",
        user="root",
        password="root"
    )

    cursor = cnx.cursor()

    query = "SELECT users.name, SUM(TIMESTAMPDIFF(SECOND, logs.start_at, IF(logs.end_at='2016-01-01 00:00:00.000', NOW(), logs.end_at))) AS total_seconds FROM users JOIN logs ON users.id = logs.user_id WHERE logs.room_id != 3 AND YEAR(logs.start_at) = YEAR(CURRENT_DATE()) AND MONTH(logs.start_at) = MONTH(CURRENT_DATE()) GROUP BY users.name"

    cursor.execute(query)

    # ユーザーごとのログ時間の合計の辞書を作成
    log_times = {}
    for row in cursor:
        log_times[row[0]] = row[1]

    cnx.close()

    return log_times
