package models

import "github.com/jinzhu/gorm"

type User struct {
    gorm.Model
    Username     string `gorm:"unique;not null"`
    PasswordHash string `gorm:"not null"`
    Roles        []Role `gorm:"many2many:user_roles;"`
}
