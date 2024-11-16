package database

import (
    "com.digitalidentity.identityservice/config"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
    dsn := config.AppConfig.DBUrl
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    return nil
}
