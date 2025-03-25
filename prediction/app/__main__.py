import uvicorn
from app.api.endpoints import probability, time
from fastapi import FastAPI


def main():
    app = FastAPI()
    app.include_router(probability.router)
    app.include_router(time.router)
    uvicorn.run(app, host="0.0.0.0", port=8085)


if __name__ == "__main__":
    main()
