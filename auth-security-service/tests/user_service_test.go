package tests

import (
    "auth-security-service/models"
    "auth-security-service/services"
    "auth-security-service/utils"
    "testing"
)

func TestCreateUser(t *testing.T) {
    db := setupTestDB()
    defer db.Close()
    userService := services.NewUserService(db)

    err := userService.CreateUser("newuser", "newpassword", []uint{})
    if err != nil {
        t.Errorf("Failed to create user: %s", err)
    }

    // Kullanıcıyı veritabanından al ve doğrula
    var user models.User
    db.Where("username = ?", "newuser").First(&user)
    if user.ID == 0 {
        t.Errorf("User not found in database")
    }

    if !utils.CheckPasswordHash("newpassword", user.PasswordHash) {
        t.Errorf("Password hash does not match")
    }
}
