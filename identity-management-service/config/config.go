package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    Port      string
    DBUrl     string
    JWTSecret string
}

var AppConfig Config

func LoadConfig() {
    viper.SetConfigFile(".env")
    viper.AutomaticEnv()

    viper.SetDefault("PORT", "8080")

    err := viper.ReadInConfig()
    if err != nil {
        // .env dosyası yoksa varsayılan değerleri kullan
    }

    AppConfig = Config{
        Port:      viper.GetString("PORT"),
        DBUrl:     viper.GetString("DATABASE_URL"),
        JWTSecret: viper.GetString("JWT_SECRET"),
    }
}
