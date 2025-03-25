from app.api.schemas import TimeResponse
from app.api.service.prediction_time import get_prediction_time
from fastapi import APIRouter, Query

router = APIRouter(
    prefix="/api/v1/prediction/time",
    tags=["time"],
    responses={404: {"description": "Not found"}},
)

@router.get(path="/", response_model=TimeResponse)
async def visit_probability(
    logs: list[str] = Query(..., description="Start at"),
    weeks: int = Query(ge=1, description="Weeks"),
) -> TimeResponse:
    t = get_prediction_time(logs, weeks)
    return TimeResponse(time=t)
