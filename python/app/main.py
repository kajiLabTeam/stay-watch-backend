import upload
import pie_chart
import data
import schedule
import time


def task():
    log_times = data.get_log_times()
    image_path = pie_chart.save_users_log_time_pie_chart(log_times)
    upload.post_slack(image_path)


# 毎月最終日の23:59にジョブを実行する
schedule.every().month.do(task).last_day().at("23:50")

while True:
    schedule.run_pending()
    time.sleep(1)
