import mlflow
from mlflow.tracking import MlflowClient
from config.settings import settings

class ModelRegistry:
    def __init__(self):
        self.client = MlflowClient(tracking_uri=settings.MODEL_REGISTRY_URI)

    def load_model(self, model_name, stage="Production"):
        model_uri = f"models:/{model_name}/{stage}"
        model = mlflow.pyfunc.load_model(model_uri)
        return model
