package repositories

import (
    "auth-security-service/models"
    "github.com/jinzhu/gorm"
)

type UserRepository struct {
    DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
    return r.DB.Create(user).Error
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
    var user models.User
    err := r.DB.Where("username = ?", username).Preload("Roles").First(&user).Error
    return &user, err
}

func (r *UserRepository) FindByID(userID uint) (*models.User, error) {
    var user models.User
    err := r.DB.Where("id = ?", userID).Preload("Roles").First(&user).Error
    return &user, err
}
