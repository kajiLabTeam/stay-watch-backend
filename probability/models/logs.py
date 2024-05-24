from __future__ import annotations
from sqlalchemy import select
from sqlalchemy.orm import Session
from . import struct as st

def get_oldest_log_by_userId(db: Session, userId) -> st.Logs:
    # logsを取得
    log: st.Logs = db.scalar(select(st.Logs).where(st.Logs.user_id == userId).order_by(st.Logs.time))
    return log