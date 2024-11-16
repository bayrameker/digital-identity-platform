package models

import "time"

type User struct {
    ID           uint      `gorm:"primaryKey"`
    Email        string    `gorm:"unique;not null" json:"email"`
    PasswordHash string    `gorm:"not null"`
    FirstName    string    `json:"first_name"`
    LastName     string    `json:"last_name"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
}

type RegisterInput struct {
    Email     string `json:"email" binding:"required,email"`
    Password  string `json:"password" binding:"required,min=8"`
    FirstName string `json:"first_name" binding:"required"`
    LastName  string `json:"last_name" binding:"required"`
}
