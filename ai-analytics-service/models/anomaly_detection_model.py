class AnomalyDetectionModel:
    def __init__(self, model_registry):
        self.model = model_registry.load_model("AnomalyDetectionModel")

    def detect(self, data):
        output = self.model.predict(data)
        return output
