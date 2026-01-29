package config

import (
	"github.com/shopspring/decimal"
)

type CreateNewCurrencyReq struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type UpdateCurrencyReq struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	IsActive bool   `json:"is_active"`
}

type CreateNewExchangeRateReq struct {
	FromCurrencyID uint            `json:"from_currency_id"`
	ToCurrencyID   uint            `json:"to_currency_id"`
	Rate           decimal.Decimal `json:"rate"`
}
