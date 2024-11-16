package models

import "time"

type MFA struct {
    ID             uint      `gorm:"primaryKey"`
    UserID         uint      `gorm:"not null;unique"`
    Secret         string    `gorm:"not null"`
    IsEnabled      bool      `gorm:"not null;default:false"`
    RecoveryCodes  string    `gorm:"type:text"`
    CreatedAt      time.Time
    UpdatedAt      time.Time
}
