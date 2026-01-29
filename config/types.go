package config

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
