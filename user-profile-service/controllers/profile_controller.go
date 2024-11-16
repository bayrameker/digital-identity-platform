package controllers

import (
    "com.digitalidentity.userprofileservice/models"
    "com.digitalidentity.userprofileservice/services"
    "net/http"
    "github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
    userID := c.Param("user_id")
    requesterID := c.GetString("user_id") // Middleware'den alınır

    profile, err := services.GetProfile(userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
        return
    }

    // Yetkilendirme kontrolü (profil sahibi mi veya izinli mi)
    if profile.UserID != requesterID {
        c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
        return
    }

    c.JSON(http.StatusOK, profile)
}

func UpdateProfile(c *gin.Context) {
    userID := c.Param("user_id")
    requesterID := c.GetString("user_id")

    if userID != requesterID {
        c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
        return
    }

    var input models.ProfileUpdateInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedProfile, err := services.UpdateProfile(userID, input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
        return
    }

    c.JSON(http.StatusOK, updatedProfile)
}

// Diğer fonksiyonlar (DeleteProfile, UploadProfilePicture vb.)
