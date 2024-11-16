from models.model_registry import ModelRegistry
from models.recommendation_model import RecommendationModel
from models.anomaly_detection_model import AnomalyDetectionModel

class ModelService:
    def __init__(self):
        self.model_registry = ModelRegistry()
        self.recommendation_model = RecommendationModel(self.model_registry)
        self.anomaly_detection_model = AnomalyDetectionModel(self.model_registry)

    def get_recommendations(self, user_features):
        predictions = self.recommendation_model.predict(user_features)
        return predictions

    def detect_anomaly(self, data):
        is_anomalous = self.anomaly_detection_model.detect(data)
        return is_anomalous
