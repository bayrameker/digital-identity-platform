from fastapi import FastAPI
from config.settings import settings
from controllers import recommendation_controller, anomaly_controller
from utils.logger import setup_logging

def create_app():
    app = FastAPI(
        title="AI Analytics Service",
        version="1.0.0",
        description="Kullanıcı davranış analizi ve öneri servisi"
    )

    # Logger'ı kur
    setup_logging()

    # Router'ları ekle
    app.include_router(recommendation_controller.router, prefix="/api/v1/recommendations", tags=["Recommendations"])
    app.include_router(anomaly_controller.router, prefix="/api/v1/anomalies", tags=["Anomalies"])

    return app

app = create_app()
