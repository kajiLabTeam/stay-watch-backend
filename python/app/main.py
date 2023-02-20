import datetime
import upload
import pie_chart
import data
import schedule
import time


def task():
    # 現在の日付を取得
    today = datetime.date.today()

    # 月の最後の日を計算
    last_day = datetime.date(today.year, today.month,
                             1) + datetime.timedelta(days=32)
    last_day = last_day.replace(day=1) - datetime.timedelta(days=1)

    # 今日が月の最後の日かどうかを判定
    if today == last_day:
        print("今日は月の最後の日です")
        log_times = data.get_log_times()
        print(log_times)
        image_path = pie_chart.save_users_log_time_pie_chart(log_times)
        upload.post_slack(image_path)


# 一日一回実行
schedule.every(1).days.do(task)

while True:
    schedule.run_pending()
    time.sleep(1)
