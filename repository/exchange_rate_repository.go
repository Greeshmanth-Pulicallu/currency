package repository

import (
	"errors"
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

func UpdateExchangeRateByID(exchangeRateId string, updateExchangeRateReq config.UpdateExchangeRateReq) error {
	result := config.DB.
		Model(&models.ExchangeRate{}).
		Where("id = ?", exchangeRateId).
		Updates(map[string]any{
			"rate": updateExchangeRateReq.Rate,
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Id does not exist")
	}
	return nil
}

func DeleteExchangeRateByID(id string) error {
	result := config.DB.
		Model(&models.ExchangeRate{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"is_active": false,
		})

	return result.Error
}

func GetExchangeRatesForPairFromDB(fromId, toId uint) models.ExchangeRate {
	var exchangeRate models.ExchangeRate
	if err := config.DB.Where("from_currency_id = ? AND to_currency_id = ?", fromId, toId).Find(&exchangeRate).Error; err != nil {
		log.Printf("Error GetAllActiveCurrenciesFromDB: %v\n", err)
		return models.ExchangeRate{}
	}

	return exchangeRate
}
