
import datetime
import upload
import chart
import data
import schedule
import time


if __name__ == '__main__':
    log_times=data.get_log_times() 

    sorted_log_times=sorted(
        log_times.items(),key=lambda x:x[1],reverse=True
    )
    first_three_keys = [t[0] for t in sorted_log_times[:3]]

    message=f'滞在時間ランキング\n1位 {first_three_keys[0]}さん\n2位 {first_three_keys[1]}さん\n3位 {first_three_keys[2]}さん\nおめでとうございます!'
    image_path=chart.save_users_log_time_bar_chart(log_times)
    upload.post_slack(image_path,message)