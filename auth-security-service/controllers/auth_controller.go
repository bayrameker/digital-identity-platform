package controllers

import (
    "auth-security-service/models"
    "auth-security-service/services"
    "auth-security-service/utils"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "net/http"
)

type AuthController struct {
    AuthService services.AuthService
}

func (ctrl *AuthController) Login(c *gin.Context) {
    var loginData struct {
        Username string `json:"username" binding:"required"`
        Password string `json:"password" binding:"required"`
    }
    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    token, err := ctrl.AuthService.AuthenticateUser(loginData.Username, loginData.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
    authService := services.NewAuthService(db)
    authController := AuthController{AuthService: authService}

    authGroup := router.Group("/auth")
    {
        authGroup.POST("/login", authController.Login)
        // Diğer auth işlemleri
    }

    // Diğer rotalar
}
