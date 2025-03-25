import numpy as np
from app.api.service.clustering import clustering
from app.api.service.date import minuts_to_time, time_to_minuts

# 1. dataをxmeansでクラスタリング
# 2. クラスタごとに中央値を取得
# 3. それぞれの中央値に重みを付けるて和を求める(クラスタのデータ数 * (クラスタのデータ数/weeks))
# 4. HH:MM形式で確率を返す

def get_prediction_time(data: list[str], weeks: int) -> str:
    data_minutes = [time_to_minuts(d) for d in data]
    # 1. dataをxmeansでクラスタリング
    if len(data_minutes) == 1:
        return minuts_to_time(data_minutes[0])
    cluster = clustering(data_minutes)
    # 2. クラスタごとに中央値を取得
    p: list[float] = []
    for c in cluster:
        p.append(c.center * (len(c.data) / len(data_minutes)))
    # 3. 確率を合計して返す
    time_minuts: float = np.sum(p)
    time_hour = minuts_to_time(time_minuts)
    return time_hour
