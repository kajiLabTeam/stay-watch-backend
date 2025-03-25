from pydantic import BaseModel


class ProbabilityResponse(BaseModel):
    probability: float

class TimeResponse(BaseModel):
    time: str

class ClusteringResult(BaseModel):
    data: list[float]
    center: float
