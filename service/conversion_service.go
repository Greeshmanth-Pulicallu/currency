package service

import (
	"errors"

	"github.com/Greeshmanth-Pulicallu/currency/config"
	"github.com/Greeshmanth-Pulicallu/currency/repository"
	"github.com/shopspring/decimal"
)

func ConvertCurrencyService(from, to, amount string) (config.CurrencyConversionRes, error) {
	var response config.CurrencyConversionRes
	fromQuery, err := repository.GetCurrencyFromDB(from)
	if err != nil {
		return response, err
	}

	toQuery, err := repository.GetCurrencyFromDB(to)
	if err != nil {
		return response, err
	}

	if len(fromQuery) == 0 || len(toQuery) == 0 {
		return response, err
	}

	exchangeRate := repository.GetExchangeRatesForPairFromDB(fromQuery[0].ID, toQuery[0].ID)

	respAmount, err := decimal.NewFromString(amount)
	if err != nil {
		return response, err
	}

	if respAmount.IsZero() || exchangeRate.Rate.IsZero() {
		return response, errors.New("zeros for amount and exchange rate")
	}

	response.From = from
	response.To = to
	response.Amount = respAmount
	response.ExchangeRate = exchangeRate.Rate
	response.ConvertedAmount = respAmount.Mul(exchangeRate.Rate)

	return response, nil
}
