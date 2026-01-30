package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	UserID string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Hash   string `gorm:"type:varchar(255);not null"`
}

func (Users) TableName() string {
	return "app_users"
}
