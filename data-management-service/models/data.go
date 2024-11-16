package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type UserData struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    UserID      string             `bson:"user_id" json:"user_id"`
    DataType    string             `bson:"data_type" json:"data_type"`
    DataContent interface{}        `bson:"data_content" json:"data_content"`
    CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
    UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type UserDataInput struct {
    DataType    string      `json:"data_type" binding:"required"`
    DataContent interface{} `json:"data_content" binding:"required"`
}
