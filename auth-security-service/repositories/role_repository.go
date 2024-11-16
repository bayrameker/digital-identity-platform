package repositories

import (
    "auth-security-service/models"
    "github.com/jinzhu/gorm"
)

type RoleRepository struct {
    DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
    return RoleRepository{DB: db}
}

func (r *RoleRepository) FindRolesByIDs(roleIDs []uint) ([]models.Role, error) {
    var roles []models.Role
    err := r.DB.Where("id IN (?)", roleIDs).Find(&roles).Error
    return roles, err
}
