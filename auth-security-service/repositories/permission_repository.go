package repositories

import (
    "auth-security-service/models"
    "github.com/jinzhu/gorm"
)

type PermissionRepository struct {
    DB *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
    return PermissionRepository{DB: db}
}

func (r *PermissionRepository) GetAllPermissions() ([]models.Permission, error) {
    var permissions []models.Permission
    err := r.DB.Find(&permissions).Error
    return permissions, err
}
