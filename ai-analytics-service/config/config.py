import os

class Config:
    KAFKA_BOOTSTRAP_SERVERS = os.getenv("KAFKA_BOOTSTRAP_SERVERS", "localhost:9092")
    KAFKA_TOPIC = os.getenv("KAFKA_TOPIC", "user_events")
    REDIS_HOST = os.getenv("REDIS_HOST", "localhost")
    REDIS_PORT = int(os.getenv("REDIS_PORT", "6379"))
    MODEL_PATH = os.getenv("MODEL_PATH", "./models/")
    SPARK_MASTER = os.getenv("SPARK_MASTER", "local[*]")
