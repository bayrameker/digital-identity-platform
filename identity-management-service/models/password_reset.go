package models

import "time"

type PasswordResetToken struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null"`
    Token     string    `gorm:"unique;not null"`
    ExpiresAt time.Time `gorm:"not null"`
    CreatedAt time.Time
}
