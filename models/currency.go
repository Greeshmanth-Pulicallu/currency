package models

import "gorm.io/gorm"

type Currency struct {
	gorm.Model
	Code     string `gorm:"uniqueIndex;not null"`
	Name     string `gorm:"not null"`
	Symbol   string `gorm:"not null"`
	IsActive bool   `gorm:"not null"`
}

func (Currency) TableName() string {
	return "app_currencies"
}
