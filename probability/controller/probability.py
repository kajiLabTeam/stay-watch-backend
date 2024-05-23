from __future__ import annotations
import math
from typing import List
from pydantic import BaseModel
from fastapi import APIRouter, Depends
from fastapi.encoders import jsonable_encoder
from fastapi.responses import ORJSONResponse
from sqlalchemy.orm import Session
from datetime import datetime, timedelta
from models import cluster as cl, user as us
from service import normal_distribution as nd
from lib.mysql import get_db

# レスポンス用のクラス
class ProbabilityResponse(BaseModel):
    userId: int
    userName: str
    probability: float

router = APIRouter()

# フロント側から現在の時刻を受け取ることとする（後で要変更)
# 今日のある時間までに特定のユーザーが入退室する確率、もしくはある時間以降に入退室する確率を返す
# 変数：user_id, true or false
@router.get("/app/probability/{reporting}/{before}" , response_class=ORJSONResponse, response_model=ProbabilityResponse)
async def get_probability_reporting_before(reporting:str, before:str, user_id:int = 0, date:str = '2024-1-1', time:str = "24:00:00", db: Session = Depends(get_db)):
    r = True if reporting == "reporting" else False
    b = True if before == "before" else False
    date_object= datetime.strptime(date, '%Y-%m-%d')
    seven_days_ago= date_object - timedelta(days=7)
    clusters = cl.get_all_cluster_by_userId_and_date(db, user_id, seven_days_ago, r)
    delta = abs(clusters[0].date - cl.get_oldest_cluster_by_userId(db, user_id, r).date + timedelta(days=1))
    # 差分を日単位に変換
    days_difference = math.floor(delta.days/7)
    # ここでクラスタリングの結果を元に確率を計算する(bがTrueなら以前, Falseなら以降)
    pr = nd.probability_from_normal_distribution(clusters, time, days_difference, b)
    result = ProbabilityResponse(userId=user_id, userName=us.get_user_by_id(db, user_id).name, probability=pr)
    result_json = jsonable_encoder(result)
    return ORJSONResponse(result_json)

# 全てのユーザがその日に入室する確率を返す
@router.get("/app/probability/{community}/all", response_class=ORJSONResponse, response_model=List[ProbabilityResponse])
async def get_probability_all(community:int, date:str = "2024-1-1", db: Session = Depends(get_db)):
    date_object= datetime.strptime(date, '%Y-%m-%d')
    seven_days_ago= date_object - timedelta(days=7)
    users = us.get_all_users_by_community(community,db)
    # 結果格納用のリスト
    result: list[ProbabilityResponse] = []
    # ユーザーごとに繰り返す
    for user in users:
        clusters = cl.get_all_cluster_by_userId_and_date(db, user.id, seven_days_ago, True)
        delta = abs(clusters[0].date - cl.get_oldest_cluster_by_userId(db, user.id, True).date + timedelta(days=1))
        # 差分を日単位に変換
        days_difference = math.floor(delta.days/7)
        # ここでクラスタリングの結果を元に確率を計算する(bがTrueなら以前, Falseなら以降)
        pr = nd.probability_from_normal_distribution(clusters, "24:00:00", days_difference, True)
        result.append(ProbabilityResponse(userId=user.id, userName=user.name, probability=pr))
    # resultをjsonに変換
    result_json = jsonable_encoder(result)
    return ORJSONResponse(result_json)