package main

import (
    "com.digitalidentity.datamanagementservice/config"
    "com.digitalidentity.datamanagementservice/controllers"
    "com.digitalidentity.datamanagementservice/database"
    "com.digitalidentity.datamanagementservice/middleware"
    "github.com/gin-gonic/gin"
)

func main() {
    // Konfigürasyonları yükle
    config.LoadConfig()

    // Veritabanına bağlan
    if err := database.Connect(); err != nil {
        panic("Failed to connect to database: " + err.Error())
    }

    // Router'ı oluştur
    router := gin.Default()

    // Middleware'leri ekle
    router.Use(middleware.JWTAuthMiddleware())

    // Endpoint'leri tanımla
    dataGroup := router.Group("/api/v1/data")
    {
        dataGroup.POST("/", controllers.CreateData)
        dataGroup.GET("/:data_id", controllers.GetData)
        dataGroup.PUT("/:data_id", controllers.UpdateData)
        dataGroup.DELETE("/:data_id", controllers.DeleteData)
        dataGroup.POST("/:data_id/permissions", controllers.GrantPermission)
        dataGroup.GET("/:data_id/logs", controllers.GetDataLogs)
    }

    // Sunucuyu başlat
    router.Run(":" + config.AppConfig.Port)
}
