import numpy as np
from datetime import date, timedelta
from datetime import datetime
from service import clustering, date_operation as do
from models import log, users, editedlog, cluster
from lib.mysql import get_db



def main():
    db = get_db().__next__()
    # user一覧を取得
    user_list = users.get_all_users(db)
    now_date = date.today()
    # 1. ユーザーごとに繰り返す
    for user in user_list:
        user_id = user.id
        # ログの取得、ログの編集、dbへの格納
        # dbからログを取得
        df = log.get_log_by_userId_and_period(db, user_id, now_date - timedelta(days=7), now_date - timedelta(days=1))
        # ログを加工
        df = do.formatting_log(df)
        # dbへ格納
        for i in range(len(df)):
            editedlog.add_edited_log(db, user_id, df['date'][i], df['reporting'][i], df['leave'][i])
        # 編集されたログから各曜日の登下校の時間をクラスタリング
        # 月〜日曜日まで7回繰り返す
        for i in range(7):
            result_entry: list[list[int]] = []
            result_exit: list[list[int]] = []
            # 1. 曜日ごとにDataFrameを取得
            df_day = editedlog.get_edited_logs_by_userId_and_day(db ,user_id, i)
            # 2. 入退室の時間をstr型のリストで取得
            df_entry: list[datetime] = df_day['reporting'].to_list()
            df_exit: list[datetime] = df_day['leaving'].to_list()
            # 時間データを秒単位に変換
            data_seconds_entry: list[int] = [sum(x * int(t) for x, t in zip([3600, 60, 1], point.strftime('%H:%M:%S').split(":"))) for point in df_entry]
            data_seconds_exit: list[int] = [sum(x * int(t) for x, t in zip([3600, 60, 1], point.strftime('%H:%M:%S').split(":"))) for point in df_exit]
            # クラスタリングの実行
            # 結果をリストに格納
            result_entry = clustering.xmeans(data_seconds_entry)
            result_exit = clustering.xmeans(data_seconds_exit)
            # 5. クラスタの数だけ繰り返す
            for j in range(len(result_entry)):
                # 入室
                # 1. クラスタの平均・標準偏差・クラスタに所属するデータの数を求める
                mean_entry = np.mean(result_entry[j])
                std_entry = np.std(result_entry[j])
                count_entry = len(result_entry[j])
                # 2. dbに格納
                day = date.today() - timedelta(days=7) + timedelta(days=i)
                cluster.add_cluster(db, user_id, day, True, mean_entry, std_entry, count_entry)
            for j in range(len(result_exit)):
                # 退室
                # 1. クラスタの平均・標準偏差・クラスタに所属するデータの数を求める
                mean_exit = np.mean(result_exit[j])
                std_exit = np.std(result_exit[j])
                count_exit = len(result_exit[j])
                # 2. dbに格納
                cluster.add_cluster(db, user_id, day, False, mean_exit, std_exit, count_exit)

if __name__ == "__main__":
    main()