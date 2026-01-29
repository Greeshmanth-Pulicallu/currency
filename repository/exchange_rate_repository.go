package repository

import (
	_ "errors"
	_ "log"

	"github.com/Greeshmanth-Pulicallu/currency/config"
	"github.com/Greeshmanth-Pulicallu/currency/models"
)

func AddNewExchangeRateToDB(exchangeRate config.CreateNewExchangeRateReq) error {
	var newExchange models.ExchangeRate

	newExchange.FromCurrencyID = exchangeRate.FromCurrencyID
	newExchange.ToCurrencyId = exchangeRate.ToCurrencyID
	newExchange.Rate = exchangeRate.Rate
	newExchange.IsActive = true

	if err := config.DB.Create(&newExchange).Error; err != nil {
		return err
	}

	return nil

}
