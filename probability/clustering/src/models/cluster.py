from sqlalchemy.orm import Session
import datetime
from . import struct as st

# クラスタをdbへ格納
def add_cluster(db: Session, uid: int, day: datetime.date, reporting: bool, mean: float, std: float, count: int) -> st.Cluster:
    # clusterを追加
    cluster:st.Cluster = st.Cluster(user_id=uid, date=day,  reporting=reporting, average=mean, sd=std, count=count)
    db.add(cluster)
    db.commit()
    db.refresh(cluster)
    return cluster