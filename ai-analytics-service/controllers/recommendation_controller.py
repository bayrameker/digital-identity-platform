from fastapi import APIRouter, Depends
from services.model_service import ModelService
from services.data_service import DataService
from utils.security import get_current_user

router = APIRouter()
model_service = ModelService()

@router.get("/", summary="Get Recommendations")
def get_recommendations(current_user: str = Depends(get_current_user)):
    # Kullanıcı için önerileri getir
    user_features = DataService.get_user_features(current_user)
    if not user_features:
        return {"error": "User features not found"}
    recommendations = model_service.get_recommendations(user_features)
    return {"user_id": current_user, "recommendations": recommendations}
