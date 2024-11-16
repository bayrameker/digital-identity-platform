package services

import (
    "com.digitalidentity.datamanagementservice/blockchain"
    "com.digitalidentity.datamanagementservice/database"
    "com.digitalidentity.datamanagementservice/models"
    "context"
    "errors"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateData(userID string, input models.UserDataInput) (models.UserData, error) {
    collection := database.MongoClient.Database("data_management").Collection("user_data")

    data := models.UserData{
        ID:          primitive.NewObjectID(),
        UserID:      userID,
        DataType:    input.DataType,
        DataContent: input.DataContent,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }

    _, err := collection.InsertOne(context.TODO(), data)
    if err != nil {
        return models.UserData{}, err
    }

    // Blok zincirine izin ekle (varsayılan olarak veri sahibine "owner" izni)
    err = blockchain.GrantPermission(data.ID.Hex(), userID, "owner")
    if err != nil {
        return models.UserData{}, err
    }

    return data, nil
}

func GetData(userID string, dataID string) (models.UserData, error) {
    // Erişim iznini kontrol et
    hasPermission, err := blockchain.CheckPermission(dataID, userID, "read")
    if err != nil || !hasPermission {
        return models.UserData{}, errors.New("access denied")
    }

    collection := database.MongoClient.Database("data_management").Collection("user_data")
    var data models.UserData

    objID, err := primitive.ObjectIDFromHex(dataID)
    if err != nil {
        return models.UserData{}, err
    }

    filter := bson.M{"_id": objID}
    err = collection.FindOne(context.TODO(), filter).Decode(&data)
    if err != nil {
        return models.UserData{}, err
    }

    // Veri erişim logunu kaydet
    LogDataAccess(userID, dataID, "read")

    return data, nil
}

// Diğer fonksiyonlar (UpdateData, DeleteData, GrantPermission, LogDataAccess)
