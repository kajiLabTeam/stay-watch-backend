from datetime import timedelta

def convert_seconds_to_hms(seconds):
    hours, remainder = divmod(seconds, 3600)
    minutes, seconds = divmod(remainder, 60)
    return f"{int(hours):02d}:{int(minutes):02d}:{int(seconds):02d}"

def number_days(start_date, end_date, weekday):
    # 開始日と終了日をdatetimeオブジェクトに変換
    # カウントを初期化
    day_count = 0
    # 開始日から終了日まで1日ずつ進みながら曜日を確認
    current_date = start_date
    while current_date <= end_date:
        # 曜日が weekday であればカウントを増やす
        if current_date.weekday() == weekday:
            day_count += 1
        # 次の日に進む
        current_date += timedelta(days=1)
    return day_count

# def first_day_of_month(date):
#     # 指定された日付をdatetimeオブジェクトに変換
#     date_obj = datetime.strptime(date.values[0][0], "%Y-%m-%d")
#     # 月の最初の日を取得
#     first_date = date_obj.replace(day=1)
#     return first_date

# def last_day_of_month(date):
#     # 指定された日付をdatetimeオブジェクトに変換
#     date_obj = datetime.strptime(date.values[0][0], "%Y-%m-%d")
#     # 月の最後の日を取得
#     next_month = date_obj.replace(day=28) + timedelta(days=4)
#     last_date = next_month - timedelta(days=next_month.day)
#     return last_date