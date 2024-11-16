package database

import (
    "com.digitalidentity.userprofileservice/config"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "context"
    "time"
)

var MongoClient *mongo.Client

func Connect() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI(config.AppConfig.DBUrl)
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return err
    }

    MongoClient = client
    return nil
}
