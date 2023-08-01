import datetime
import upload
import chart
import data
import schedule
import time


def task():
    # 現在時刻を表示
    print("Function executed at:", datetime.datetime.now())
    # 現在の日付を取得
    today = datetime.date.today()

    # 今月の最初の日を計算
    first_day = datetime.date(today.year, today.month, 1)

    # 今日が月の最初の日かどうかを判定
    if today == first_day:
        log_times = data.get_log_times()

        sorted_log_times=sorted(
            log_times.items(),key=lambda x:x[1],reverse=True
        )
        first_three_keys = [t[0] for t in sorted_log_times[:3]]
        message=f'滞在時間ランキング\n1位 {first_three_keys[0]}さん\n2位 {first_three_keys[1]}さん\n3位 {first_three_keys[2]}さん\nおめでとうございます!'

        image_path = chart.save_users_log_time_bar_chart(log_times,message)
        upload.post_slack(image_path,message) 

# 一日一回実行
schedule.every(1).days.do(task)

while True:
    schedule.run_pending()
    time.sleep(1)
