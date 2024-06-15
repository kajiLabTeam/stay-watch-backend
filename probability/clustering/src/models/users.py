from __future__ import annotations
from sqlalchemy.orm import Session
from sqlalchemy import select
from . import struct as st

# userを全て取得
def get_all_users(db: Session) -> list[st.User]:
    # usersを取得
    users = db.scalars(select(st.User)).all()
    return users