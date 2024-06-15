# sqlalchemyライブラリから使用する型などをインポート
from sqlalchemy.orm import (
    DeclarativeBase,
    Mapped,
    mapped_column
)
import datetime

# Baseクラスを作成
class Base(DeclarativeBase):
    pass

# Baseクラスを継承したモデルを作成
# # usersテーブルのモデルUsers
# class Users(Base):
#     __tablename__ = 'users'
#     user_id = mapped_column(Integer, primary_key=True, autoincrement=True)
#     uid = mapped_column(String(255), nullable=False)
#     name = mapped_column(String(255), nullable=False)
#     email = mapped_column(String(255), nullable=False)
#     role = mapped_column(String(255), nullable=False)
# logs(仮)テーブルのモデルLogs(仮)
class EditedLog(Base):
    __tablename__ = 'edited_logs'
    id: Mapped[int] = mapped_column(primary_key=True, index=True)
    user_id: Mapped[int] = mapped_column(nullable=False)
    date: Mapped[datetime.date] = mapped_column(nullable=False)
    reporting: Mapped[datetime.time] = mapped_column(nullable=False)
    leaving: Mapped[datetime.time] = mapped_column(nullable=False)
# clusterテーブルのモデルCluster
class Cluster(Base):
    __tablename__ = 'clusters'
    id: Mapped[int] = mapped_column(primary_key=True, autoincrement=True)
    date: Mapped[datetime.date] = mapped_column(nullable=False)
    reporting: Mapped[bool] = mapped_column(nullable=False)
    average: Mapped[float] = mapped_column(nullable=False)
    sd: Mapped[float] = mapped_column(nullable=False)
    count: Mapped[int] = mapped_column(nullable=False)
    user_id: Mapped[int] = mapped_column(nullable=False)

class Log(Base):
    __tablename__ = 'logs'
    id: Mapped[int] = mapped_column(primary_key=True, autoincrement=True)
    created_at: Mapped[datetime.datetime] = mapped_column(nullable=False)
    updated_at: Mapped[datetime.datetime] = mapped_column(nullable=False)
    deleted_at: Mapped[datetime.datetime] = mapped_column(nullable=True)
    room_id: Mapped[int] = mapped_column(nullable=False)
    start_at: Mapped[datetime.datetime] = mapped_column(nullable=False)
    end_at: Mapped[datetime.datetime] = mapped_column(nullable=False)
    user_id: Mapped[int] = mapped_column(nullable=False)
    rssi: Mapped[int] = mapped_column(nullable=False)

class User(Base):
    __tablename__ = 'users'
    id: Mapped[int] = mapped_column(primary_key=True, autoincrement=True)
    created_at: Mapped[datetime.datetime] = mapped_column(nullable=False)
    updated_at: Mapped[datetime.datetime] = mapped_column(nullable=False)
    deleted_at: Mapped[datetime.datetime] = mapped_column(nullable=True)
    uuid: Mapped[str] = mapped_column(nullable=False)
    name: Mapped[str] = mapped_column(nullable=False)
    email: Mapped[str] = mapped_column(nullable=False)
    role: Mapped[int] = mapped_column(nullable=False)
    beacon_id: Mapped[int] = mapped_column(nullable=False)
    community_id: Mapped[int] = mapped_column(nullable=False)