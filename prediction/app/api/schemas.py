from pydantic import BaseModel


class ProbabilityResponse(BaseModel):
    probability: float


class ClusteringResult(BaseModel):
    data: list[float]
    center: float
