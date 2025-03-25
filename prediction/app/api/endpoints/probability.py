from app.api.schemas import ProbabilityResponse
from app.api.service.probability import get_probability
from fastapi import APIRouter, Query

router = APIRouter(
    prefix="/api/v1/prediction/probability",
    tags=["probability"],
    responses={404: {"description": "Not found"}},
)


@router.get(path="/", response_model=ProbabilityResponse)
async def visit_probability(
    logs: list[str] = Query(..., description="Start at"),
    time: str = Query(..., regex=r"^\d{2}:\d{2}$", description="Time"),
    weeks: int = Query(ge=1, description="Weeks"),
    is_forward: bool = Query(True, description="Is forward"),
) -> ProbabilityResponse:
    p = get_probability(logs, time, weeks)
    if not is_forward:
        p = 1 * (len(logs) / weeks) - p
    return ProbabilityResponse(probability=p)
