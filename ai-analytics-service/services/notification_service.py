import requests

class NotificationService:
    @staticmethod
    def send_notification(user_id, message):
        payload = {"user_id": user_id, "message": message}
        response = requests.post(f"http://notification-service/api/v1/notifications", json=payload)
        return response.status_code == 200
