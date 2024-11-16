package models

import "github.com/jinzhu/gorm"

type Permission struct {
    gorm.Model
    Name string `gorm:"unique;not null"`
}
