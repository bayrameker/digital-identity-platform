package main

import (
    "com.digitalidentity.identityservice/controllers"
    "com.digitalidentity.identityservice/middleware"
    "com.digitalidentity.identityservice/config"
    "com.digitalidentity.identityservice/database"
    "github.com/gin-gonic/gin"
)

func main() {
    // Konfigürasyonları yükle
    config.LoadConfig()

    // Veritabanına bağlan
    err := database.Connect()
    if err != nil {
        panic("Failed to connect to database!")
    }

    // Router'ı oluştur
    r := gin.Default()

    // Middleware'leri ekle
    r.Use(middleware.LoggingMiddleware())

    // Endpoint'leri tanımla
    auth := r.Group("/api/v1/auth")
    {
        auth.POST("/register", controllers.Register)
        auth.POST("/login", controllers.Login)
        auth.POST("/logout", middleware.JWTAuthMiddleware(), controllers.Logout)
        auth.POST("/refresh", controllers.RefreshToken)
        auth.POST("/password/reset", controllers.RequestPasswordReset)
        auth.POST("/password/reset/confirm", controllers.ResetPassword)
    }

    // Diğer endpoint grupları (kullanıcı profili, MFA vb.)

    // Sunucuyu başlat
    r.Run(":" + config.AppConfig.Port)
}
