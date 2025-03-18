from typing import Optional

from app.api.schemas import ProbabilityResponse
from app.api.service.probability import get_probability
from fastapi import APIRouter, Query

router = APIRouter(
    prefix="/api/v1/prediction/probability",
    tags=["probability"],
    responses={404: {"description": "Not found"}},
)


@router.get("/visit", response_model=ProbabilityResponse)
async def visit_probability(
    start_at: list[str] = Query(..., description="Start at"),
    time: str = Query(..., regex=r"^\d{2}:\d{2}$", description="Time"),
    weeks: int = Query(ge=1, description="Weeks"),
    is_forward: bool = Query(True, description="Is forward"),
) -> ProbabilityResponse:
    p = get_probability(start_at, time, weeks)
    if not is_forward:
        p = 1 * (len(start_at) / weeks) - p
    return ProbabilityResponse(probability=p)


@router.get("/departure", response_model=ProbabilityResponse)
async def departure_probability(
    end_at: list[str] = Query(..., description="End at"),
    time: str = Query(..., regex=r"^\d{2}:\d{2}$", description="Time"),
    weeks: int = Query(ge=1, description="Weeks"),
    is_forward: Optional[bool] = Query(True, description="Is forward"),
) -> ProbabilityResponse:
    p = get_probability(end_at, time, weeks)
    if not is_forward:
        p = 1 * (len(end_at) / weeks) - p
    return ProbabilityResponse(probability=p)
