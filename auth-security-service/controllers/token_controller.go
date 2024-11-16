package controllers

import (
    "auth-security-service/services"
    "github.com/gin-gonic/gin"
    "net/http"
)

type TokenController struct {
    TokenService services.TokenService
}

func (ctrl *TokenController) RefreshToken(c *gin.Context) {
    var requestData struct {
        RefreshToken string `json:"refresh_token" binding:"required"`
    }
    if err := c.ShouldBindJSON(&requestData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    newToken, err := ctrl.TokenService.RefreshAccessToken(requestData.RefreshToken)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": newToken})
}
