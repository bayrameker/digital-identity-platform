package main

import (
    "auth-security-service/config"
    "auth-security-service/controllers"
    "auth-security-service/middlewares"
    "auth-security-service/utils"
    "github.com/gin-gonic/gin"
)

func main() {
    // Konfigürasyonu yükle
    config.LoadConfig()

    // Logger'ı başlat
    utils.SetupLogger()

    // Veritabanına bağlan
    db := config.ConnectDatabase()
    defer db.Close()

    // Gin motorunu oluştur
    router := gin.Default()

    // Middleware'leri ekle
    router.Use(middlewares.LoggingMiddleware())
    router.Use(middlewares.CORSMiddleware())

    // Rotaları tanımla
    controllers.RegisterRoutes(router, db)

    // Sunucuyu başlat
    router.RunTLS(config.AppConfig.ServerAddress, config.AppConfig.TLSCertFile, config.AppConfig.TLSKeyFile)
}
