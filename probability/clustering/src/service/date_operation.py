import pandas as pd
from datetime import datetime

def identification_day(df: pd.DataFrame) -> pd.DataFrame:
    df_date= df['date'].to_list()
    # 日付から曜日を取得してDataFrameに追加
    day: list[int] = []
    for i in range(len(df_date)):
        # 時間の文字列をdatetime型に変換
        date_object = datetime.strptime(df_date[i], '%Y-%m-%d')
        day.append(date_object.weekday())
    df['day'] = day
    # print(df)
    return df


def formatting_log(df: pd.DataFrame) -> pd.DataFrame:
    # 時間の文字列をdatetime型に変換
    df['start_at'] = pd.to_datetime(df['start_at'])
    df['end_at'] = pd.to_datetime(df['end_at'])

    # 日付ごとに最初の入室と最後の入室を取得
    result_data = []
    grouped_data = df.groupby(df['start_at'].dt.date)
    for date, group in grouped_data:
        data = group.loc[group['start_at'].idxmin()]
        result_data.append({
            'user_id': df['user_id'],
            'date': date,
            'reporting': data['start_at'].strftime('%H:%M:%S.%f')[:-3],
            'leave': data['start_at'].strftime('%H:%M:%S.%f')[:-3],
        })

    # 結果をDataFrameに変換してreturn
    result_df = pd.DataFrame(result_data)
    return result_df