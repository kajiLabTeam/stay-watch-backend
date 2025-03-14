import uvicorn
from app.api.endpoints import probability
from fastapi import FastAPI


def main():
    app = FastAPI(
        prefix="/api/v1/prediction", responses={404: {"description": "Not found"}}
    )
    app.include_router(probability.router)
    uvicorn.run(app, host="0.0.0.0", port=8085)


if __name__ == "__main__":
    main()
