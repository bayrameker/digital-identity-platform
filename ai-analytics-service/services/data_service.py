import requests
from config.settings import settings

class DataService:
    @staticmethod
    def get_user_features(user_id):
        # Veri Yönetimi Servisi'nden kullanıcı özelliklerini al
        response = requests.get(f"http://data-management-service/api/v1/users/{user_id}/features")
        if response.status_code == 200:
            return response.json()
        else:
            return None

    @staticmethod
    def save_recommendations(user_id, recommendations):
        # Önerileri kaydet veya başka bir servise ilet
        payload = {"user_id": user_id, "recommendations": recommendations}
        response = requests.post(f"http://recommendation-service/api/v1/recommendations", json=payload)
        return response.status_code == 200
