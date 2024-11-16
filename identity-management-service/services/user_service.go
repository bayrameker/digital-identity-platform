package services

import (
    "com.digitalidentity.identityservice/models"
    "com.digitalidentity.identityservice/database"
    "com.digitalidentity.identityservice/utils"
    "errors"
)

func CreateUser(input models.RegisterInput) (models.User, error) {
    // Parolayı hashle
    hashedPassword, err := utils.HashPassword(input.Password)
    if err != nil {
        return models.User{}, err
    }

    // Kullanıcı modelini oluştur
    user := models.User{
        Email:        input.Email,
        PasswordHash: hashedPassword,
        FirstName:    input.FirstName,
        LastName:     input.LastName,
    }

    // Veritabanına kaydet
    result := database.DB.Create(&user)
    if result.Error != nil {
        return models.User{}, result.Error
    }

    return user, nil
}

// Diğer fonksiyonlar (AuthenticateUser, GetUserByID vb.)
