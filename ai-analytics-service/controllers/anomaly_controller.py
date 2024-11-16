from fastapi import APIRouter, Depends
from pydantic import BaseModel
from services.model_service import ModelService
from utils.security import get_current_user

router = APIRouter()
model_service = ModelService()

class AnomalyDetectionRequest(BaseModel):
    data: dict

@router.post("/detect", summary="Detect Anomaly")
def detect_anomaly(request: AnomalyDetectionRequest, current_user: str = Depends(get_current_user)):
    is_anomalous = model_service.detect_anomaly(request.data)
    return {"user_id": current_user, "is_anomalous": is_anomalous}
