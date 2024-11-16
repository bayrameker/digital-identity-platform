from kafka import KafkaConsumer
from services.model_service import ModelService
from services.data_service import DataService
from services.notification_service import NotificationService
from config.settings import settings
from utils.logger import logger
import json

class KafkaConsumerService:
    def __init__(self):
        self.consumer = KafkaConsumer(
            settings.KAFKA_USER_EVENTS_TOPIC,
            bootstrap_servers=settings.KAFKA_BOOTSTRAP_SERVERS,
            value_deserializer=lambda x: json.loads(x.decode('utf-8')),
            auto_offset_reset='latest',
            enable_auto_commit=True,
            group_id='ai-analytics-group'
        )
        self.model_service = ModelService()

    def consume(self):
        for message in self.consumer:
            event = message.value
            self.process_event(event)

    def process_event(self, event):
        user_id = event.get('user_id')
        user_features = event.get('features')
        if not user_id or not user_features:
            logger.error("Invalid event data")
            return

        # Anomali tespiti
        is_anomalous = self.model_service.detect_anomaly(user_features)
        if is_anomalous:
            # Anomali tespit edildi, bildirim gönder
            NotificationService.send_notification(user_id, "Anormal etkinlik tespit edildi.")
            logger.info(f"Anomaly detected for user {user_id}")

        # Öneri oluşturma
        recommendations = self.model_service.get_recommendations(user_features)
        # Önerileri kaydet veya gönder
        success = DataService.save_recommendations(user_id, recommendations)
        if success:
            logger.info(f"Recommendations saved for user {user_id}")
        else:
            logger.error(f"Failed to save recommendations for user {user_id}")
