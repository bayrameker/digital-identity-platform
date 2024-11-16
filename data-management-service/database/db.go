package database

import (
    "com.digitalidentity.datamanagementservice/config"
    "context"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func Connect() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI(config.AppConfig.MongoDBURI)
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return err
    }

    // Bağlantıyı doğrula
    err = client.Ping(ctx, nil)
    if err != nil {
        return err
    }

    MongoClient = client
    return nil
}
