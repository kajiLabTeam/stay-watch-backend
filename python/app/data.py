
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

    # SQLクエリを部分ごとに分ける
    query_select = "SELECT users.name, SUM(TIMESTAMPDIFF(SECOND, logs.start_at, IF(logs.end_at='2016-01-01 00:00:00.000', NOW(), logs.end_at))) AS total_seconds"
    query_from = "FROM users JOIN logs ON users.id = logs.user_id"
    query_where = "WHERE logs.room_id != 3 AND users.name NOT IN ('108B(rui)', '108B(togawa)','108B(10000)')"
    query_time = "AND YEAR(logs.start_at) = YEAR(DATE_SUB(CURRENT_DATE(), INTERVAL 1 MONTH)) AND MONTH(logs.start_at) = MONTH(DATE_SUB(CURRENT_DATE(), INTERVAL 1 MONTH))"
    query_group = "GROUP BY users.name"

    # 各部分を結合して最終的なクエリを作る
    query = f"{query_select} {query_from} {query_where} {query_time} {query_group};"

    cursor.execute(query)

    # ユーザーごとのログ時間の合計の辞書を作成
    log_times = {}
    for row in cursor:
        total_seconds=row[1]
        hours, remainder = divmod(total_seconds, 3600)      # 1時間は3600秒
        minutes = remainder // 60  
        log_times[row[0]] = int(hours) + round((minutes / 60),1)

    cnx.close()

    return log_times
