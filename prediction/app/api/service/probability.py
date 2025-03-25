import numpy as np
from app.api.service.clustering import clustering
from app.api.service.date import time_to_minuts
from scipy.stats import norm


# 1. dataをxmeansでクラスタリング
# 2. クラスタごとに確率を計算
# 2-1. クラスタの中心をクラスタの平均とする
# 2-2. クラスタの標準偏差を求める
# 2-3. 正規分布の確率密度関数を用いて確率を計算
# 2-4. 算出した確率に重みをつける(クラスタのデータ数 * (クラスタのデータ数/weeks))
# 3. 確率を合計して返す
def get_probability(data: list[str], time: str, weeks: int) -> float:
    data_minutes = [time_to_minuts(d) for d in data]
    time_minutes = time_to_minuts(time)
    # 1. dataをxmeansでクラスタリング
    if len(data_minutes) == 1:
        if time_minutes >= data_minutes[0]:
            return 1 / weeks
        else:
            return 0
    cluster = clustering(data_minutes)
    # 2. クラスタごとに確率を計算
    p: list[float] = []
    for c in cluster:
        # クラスタに所属するデータが1つの場合
        if len(c.data) == 1:
            if time_minutes >= c.data[0]:
                p.append(1 / weeks)
            else:
                p.append(0)
            continue
        # 2-1. クラスタの中心をクラスタの平均とする
        loc = c.center
        # 2-2. クラスタの標準偏差を求める
        scale = float(np.std(c.data))
        # scale = 0の場合(クラスタのデータが全て同じ場合)
        if scale == 0:
            if cluster[0] == loc and time_minutes >= loc:
                p.append(1 * (len(c.data) / weeks))
            else:
                p.append(0)
            continue
        # 2-3. 正規分布の確率密度関数を用いて確率を計算
        p.append(float(norm.cdf(time_minutes, loc, scale)) * (len(c.data) / weeks))
    # 3. 確率を合計して返す
    probability = np.sum(p)
    return probability
