class RecommendationModel:
    def __init__(self, model_registry):
        self.model = model_registry.load_model("RecommendationModel")

    def predict(self, user_features):
        predictions = self.model.predict(user_features)
        return predictions
