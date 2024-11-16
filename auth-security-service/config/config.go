package config

import (
    "fmt"
    "log"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/spf13/viper"
)

type Config struct {
    ServerAddress string
    TLSCertFile   string
    TLSKeyFile    string
    Database      struct {
        Host     string
        Port     int
        User     string
        Password string
        Name     string
        SSLMode  string
    }
    JWTSecretKey string
}

var AppConfig Config

func LoadConfig() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Config file error: %s", err)
    }
    if err := viper.Unmarshal(&AppConfig); err != nil {
        log.Fatalf("Config unmarshal error: %s", err)
    }
}

func ConnectDatabase() *gorm.DB {
    dbConfig := AppConfig.Database
    connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
        dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Name, dbConfig.Password, dbConfig.SSLMode)
    db, err := gorm.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Database connection error: %s", err)
    }
    return db
}
