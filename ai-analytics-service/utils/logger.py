from loguru import logger
from config.settings import settings

def setup_logging():
    logger.remove()
    logger.add(
        "logs/app.log",
        rotation="1 MB",
        retention="7 days",
        compression="zip",
        level=settings.LOG_LEVEL,
        format="{time} {level} {message}",
    )
