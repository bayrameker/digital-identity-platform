package models

import "time"

type Profile struct {
    ID             uint      `gorm:"primaryKey" json:"id"`
    UserID         uint      `gorm:"not null" json:"user_id"`
    FirstName      string    `json:"first_name"`
    LastName       string    `json:"last_name"`
    Bio            string    `json:"bio,omitempty"`
    ProfilePicture string    `json:"profile_picture,omitempty"`
    Preferences    string    `json:"preferences,omitempty"` // JSON formatında saklanabilir
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}
