package tests

import (
    "auth-security-service/config"
    "auth-security-service/models"
    "auth-security-service/services"
    "auth-security-service/utils"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "testing"
)

func setupTestDB() *gorm.DB {
    dbConfig := config.AppConfig.Database
    connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
        dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Name+"_test", dbConfig.Password, dbConfig.SSLMode)
    db, err := gorm.Open("postgres", connStr)
    if err != nil {
        panic("Failed to connect to test database")
    }
    db.AutoMigrate(&models.User{})
    return db
}

func TestAuthenticateUser(t *testing.T) {
    db := setupTestDB()
    defer db.Close()
    authService := services.NewAuthService(db)

    // Test kullanıcı oluştur
    password := "testpassword"
    hashedPassword, _ := utils.HashPassword(password)
    user := models.User{
        Username:     "testuser",
        PasswordHash: hashedPassword,
    }
    db.Create(&user)

    // Doğru şifre ile kimlik doğrulama
    token, err := authService.AuthenticateUser("testuser", "testpassword")
    if err != nil || token == "" {
        t.Errorf("Authentication failed with correct credentials")
    }

    // Yanlış şifre ile kimlik doğrulama
    _, err = authService.AuthenticateUser("testuser", "wrongpassword")
    if err == nil {
        t.Errorf("Authentication succeeded with incorrect credentials")
    }
}
