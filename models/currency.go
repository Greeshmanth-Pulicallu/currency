package models

import "gorm.io/gorm"

type Currency struct {
	gorm.Model `json:"-"`
	Code       string `gorm:"uniqueIndex;not null" json:"code"`
	Name       string `gorm:"not null" json:"name"`
	Symbol     string `gorm:"not null" json:"symbol"`
	IsActive   bool   `gorm:"not null" json:"is_active"`
}

func (Currency) TableName() string {
	return "app_currencies"
}
