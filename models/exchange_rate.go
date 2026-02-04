package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ExchangeRate struct {
	gorm.Model `json:"-"`

	FromCurrencyID uint     `gorm:"not null;uniqueIndex:idx_from_to" json:"from_currency_id"`
	FromCurrency   Currency `gorm:"foreignKey:FromCurrencyID;references:ID" json:"from_currency"`
	ToCurrencyId   uint     `gorm:"not null;uniqueIndex:idx_from_to" json:"to_currency_id"`
	ToCurrency     Currency `gorm:"foreignKey:ToCurrencyID;references:ID" json:"to_currency"`

	Rate     decimal.Decimal `gorm:"type:numeric(18,6);not null" json:"rate"`
	IsActive bool            `gorm:"not null" json:"is_active"`
}

func (ExchangeRate) TableName() string {
	return "app_exchange_rates"
}
