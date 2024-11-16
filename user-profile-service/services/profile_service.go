package services

import (
    "com.digitalidentity.userprofileservice/models"
    "com.digitalidentity.userprofileservice/database"
    "context"
    "time"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProfile(userID string) (models.Profile, error) {
    var profile models.Profile
    collection := database.MongoClient.Database("user_profiles").Collection("profiles")

    filter := bson.M{"user_id": userID}
    err := collection.FindOne(context.TODO(), filter).Decode(&profile)
    if err != nil {
        return models.Profile{}, err
    }

    return profile, nil
}

func UpdateProfile(userID string, input models.ProfileUpdateInput) (models.Profile, error) {
    collection := database.MongoClient.Database("user_profiles").Collection("profiles")

    update := bson.M{
        "$set": bson.M{
            "first_name":  input.FirstName,
            "last_name":   input.LastName,
            "bio":         input.Bio,
            "preferences": input.Preferences,
            "updated_at":  time.Now(),
        },
    }

    filter := bson.M{"user_id": userID}
    _, err := collection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return models.Profile{}, err
    }

    return GetProfile(userID)
}

// Diğer fonksiyonlar (DeleteProfile, UploadProfilePicture vb.)
