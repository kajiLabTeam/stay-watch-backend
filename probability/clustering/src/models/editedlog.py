import pandas as pd
from sqlalchemy.orm import Session
from . import struct as st

# userIdから特定の曜日のEditedLogを取得し、dataframeで返す
def get_edited_logs_by_userId_and_day(db: Session, userId: int, day: int) -> pd.DataFrame:
    # edited_logsを取得
    q = db.query(st.EditedLog).filter(st.EditedLog.user_id == userId)
    edited_logs = pd.read_sql(q.statement, db.bind)
    # 指定の曜日のデータのみを抽出
    date = pd.to_datetime(edited_logs['date'])
    edited_logs['day'] = date.dt.weekday
    edited_logs = edited_logs[edited_logs['day'] == day]
    return edited_logs

# EditedLogを追加する
def add_edited_log(db: Session, uid: int, date: str, reporting_time: str, leave_time: str) -> st.EditedLog:
    # edited_logを追加
    edited_log:st.EditedLog = st.EditedLog(user_id=uid, date=date, reporting=reporting_time, leave=leave_time)
    db.add(edited_log)
    db.commit()
    db.refresh(edited_log)
    return edited_log