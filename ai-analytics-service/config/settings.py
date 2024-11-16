import os
from pydantic import BaseSettings

class Settings(BaseSettings):
    # Genel Ayarlar
    APP_NAME: str = "AI Analytics Service"
    VERSION: str = "1.0.0"
    DESCRIPTION: str = "Kullanıcı davranış analizi ve öneri servisi"

    # Kafka Ayarları
    KAFKA_BOOTSTRAP_SERVERS: str = os.getenv("KAFKA_BOOTSTRAP_SERVERS", "localhost:9092")
    KAFKA_USER_EVENTS_TOPIC: str = os.getenv("KAFKA_USER_EVENTS_TOPIC", "user_events")

    # Redis Ayarları
    REDIS_HOST: str = os.getenv("REDIS_HOST", "localhost")
    REDIS_PORT: int = int(os.getenv("REDIS_PORT", "6379"))

    # Model Ayarları
    MODEL_REGISTRY_URI: str = os.getenv("MODEL_REGISTRY_URI", "./models/")

    # Spark Ayarları
    SPARK_MASTER: str = os.getenv("SPARK_MASTER", "local[*]")

    # Güvenlik Ayarları
    JWT_SECRET_KEY: str = os.getenv("JWT_SECRET_KEY", "your_jwt_secret_key")
    JWT_ALGORITHM: str = "HS256"

    # Diğer Ayarlar
    LOG_LEVEL: str = os.getenv("LOG_LEVEL", "INFO")

settings = Settings()
