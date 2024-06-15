from __future__ import annotations
import pandas as pd
from datetime import date
from sqlalchemy.orm import Session
from . import struct as st

# userIdから指定期間のlogを取得し、dataframeで返す
def get_log_by_userId_and_period(db: Session, userId: int, start: date, end: date) -> pd.DataFrame:
    # Logを取得
    q = db.query(st.Log.start_at, st.Log.end_at, st.Log.user_id).filter(st.Log.user_id == userId, st.Log.start_at >= start, st.Log.start_at<= end)
    logs = pd.read_sql(q.statement, db.bind)
    return logs

# userIdから全てのlogを取得し、dataframeで返す
def get_all_logs_by_userId(db: Session, userId: int) -> pd.DataFrame:
    # Logを取得
    q = db.query(st.Log.start_at, st.Log.end_at, st.Log.user_id).filter(st.Log.user_id == userId)
    logs = pd.read_sql(q.statement, db.bind)
    return logs