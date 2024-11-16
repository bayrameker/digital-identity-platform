package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    Port             string
    MongoDBURI       string
    JWTSecret        string
    EthereumEndpoint string
    ContractAddress  string
}

var AppConfig Config

func LoadConfig() {
    viper.SetConfigFile(".env")
    viper.AutomaticEnv()

    viper.SetDefault("PORT", "8082")

    err := viper.ReadInConfig()
    if err != nil {
        // .env dosyası yoksa varsayılan değerleri kullan
    }

    AppConfig = Config{
        Port:             viper.GetString("PORT"),
        MongoDBURI:       viper.GetString("MONGODB_URI"),
        JWTSecret:        viper.GetString("JWT_SECRET"),
        EthereumEndpoint: viper.GetString("ETHEREUM_ENDPOINT"),
        ContractAddress:  viper.GetString("CONTRACT_ADDRESS"),
    }
}
