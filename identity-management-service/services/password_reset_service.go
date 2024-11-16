package services

import (
	"time"

	"com.digigtalidentity.identityservice/database"
	"com.digigtalidentity.identityservice/models"
	"com.digigtalidentity.identityservice/utils"
)

func CreatePasswordResetToken(userID uint) (models.PasswordResetToken, error) {
	tokenString := utils.GenerateRandomString(32)
	token := models.PasswordResetToken{
		UserID:    userID,
		Token:     tokenString,
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}

	result := database.DB.Create(&token)
	if result.Error != nil {
		return models.PasswordResetToken{}, result.Error
	}

	return token, nil
}

// Diğer fonksiyonlar (VerifyPasswordResetToken, ResetPassword vb.)
