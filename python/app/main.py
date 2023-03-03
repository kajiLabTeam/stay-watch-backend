import datetime
import upload
import pie_chart
import data
import schedule
import time


# def task():
#     # 現在時刻を表示
#     print("Function executed at:", datetime.datetime.now())
#     # 現在の日付を取得
#     today = datetime.date.today()

#     # 今月の最初の日を計算
#     first_day = datetime.date(today.year, today.month, 1)

#     # 今日が月の最初の日かどうかを判定
#     if today == first_day:
#         print("今日は月の最初の日です")
#         log_times = data.get_log_times()
#         print(log_times)
#         image_path = pie_chart.save_users_log_time_pie_chart(log_times)
#         upload.post_slack(image_path)


# # 一日一回実行
# schedule.every(1).days.do(task)

# while True:
#     schedule.run_pending()
#     time.sleep(1)


def task():
    log_times = data.get_log_times()
    print(log_times)
    image_path = pie_chart.save_users_log_time_pie_chart(log_times)
    upload.post_slack(image_path)


if __name__ == "__main__":
    task()
