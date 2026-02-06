package dto

import (
	"github.com/shopspring/decimal"
)

type CreateNewCurrencyReq struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type UpdateCurrencyReq struct {
	Name     *string `json:"name"`
	Symbol   *string `json:"symbol"`
	IsActive *bool   `json:"is_active"`
}

type CreateNewExchangeRateReq struct {
	FromCurrencyID uint            `json:"from_currency_id"`
	ToCurrencyID   uint            `json:"to_currency_id"`
	Rate           decimal.Decimal `json:"rate"`
}

type UpdateExchangeRateReq struct {
	Rate decimal.Decimal `json:"rate"`
}

type CurrencyConversionRes struct {
	From            string          `json:"from"`
	To              string          `json:"to"`
	Amount          decimal.Decimal `json:"amount"`
	ExchangeRate    decimal.Decimal `json:"exchange_rate"`
	ConvertedAmount decimal.Decimal `json:"converted_amount"`
}

type RegisterReq struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

type LoginReq struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}
