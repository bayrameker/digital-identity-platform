package services

import (
	"com.digigtalidentity.identityservice/database"
	"com.digigtalidentity.identityservice/models"
	"github.com/pquerna/otp/totp"
)

func GenerateMFASecret(userID uint) (string, error) {
	secret := totp.GenerateOpts{
		Issuer:      "YourAppName",
		AccountName: "user@example.com",
	}

	key, err := totp.Generate(secret)
	if err != nil {
		return "", err
	}

	mfa := models.MFA{
		UserID:    userID,
		Secret:    key.Secret(),
		IsEnabled: false,
	}

	result := database.DB.Create(&mfa)
	if result.Error != nil {
		return "", result.Error
	}

	return key.URL(), nil
}

func ValidateMFAToken(userID uint, token string) (bool, error) {
	var mfa models.MFA
	result := database.DB.Where("user_id = ?", userID).First(&mfa)
	if result.Error != nil {
		return false, result.Error
	}

	valid := totp.Validate(token, mfa.Secret)
	return valid, nil
}
