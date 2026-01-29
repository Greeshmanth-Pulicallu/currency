package repository

import (
	_ "errors"
	"log"

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

func GetAllActiveExchangesFromDB() ([]models.ExchangeRate, error) {
	var exchanges []models.ExchangeRate
	if err := config.DB.Where("is_active = ?", true).Find(&exchanges).Error; err != nil {
		log.Printf("Error GetAllActiveCurrenciesFromDB: %v\n", err)
		return []models.ExchangeRate{}, err
	}

	return exchanges, nil
}

func GetExchangeRatesByIDFromDB(exchangeRateId string) (models.ExchangeRate, error) {
	var currency models.ExchangeRate
	if err := config.DB.Where("id = ?", exchangeRateId).First(&currency).Error; err != nil {
		return models.ExchangeRate{}, err
	}
	return currency, nil

}
