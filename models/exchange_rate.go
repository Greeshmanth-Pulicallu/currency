package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ExchangeRate struct {
	gorm.Model

	FromCurrencyID uint     `gorm:"not null;uniqueIndex:idx_from_to"`
	FromCurrency   Currency `gorm:"foreignKey:FromCurrencyID;references:ID"`

	ToCurrencyId uint     `gorm:"not null;uniqueIndex:idx_from_to"`
	ToCurrency   Currency `gorm:"foreignKey:ToCurrencyID;references:ID"`

	Rate     decimal.Decimal `gorm:"type:numeric(18,6);not null"`
	IsActive bool            `gorm:"not null"`
}

func (ExchangeRate) TableName() string {
	return "app_exchange_rates"
}
