package repository

import (
	"errors"
	"log"

	"github.com/Greeshmanth-Pulicallu/currency/config"
	"github.com/Greeshmanth-Pulicallu/currency/dto"
	"github.com/Greeshmanth-Pulicallu/currency/models"
	"gorm.io/gorm/clause"
)

const ExchngeRateURL = "https://v6.exchangerate-api.com/v6/8759116ef8e1f7321d5eba63/latest/"

func AddNewExchangeRateToDB(exchangeRate dto.CreateNewExchangeRateReq) error {
	newExchange := models.ExchangeRate{
		FromCurrencyID: exchangeRate.FromCurrencyID,
		ToCurrencyId:   exchangeRate.ToCurrencyID,
		Rate:           exchangeRate.Rate,
		IsActive:       true,
	}

	return config.DB.
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "from_currency_id"},
				{Name: "to_currency_id"},
			},
			DoUpdates: clause.AssignmentColumns([]string{
				"rate",
				"is_active",
				"updated_at",
			}),
		}).
		Create(&newExchange).Error
}

func GetAllActiveExchangesFromDB() ([]models.ExchangeRate, error) {
	var exchanges []models.ExchangeRate
	if err := config.DB.Where("is_active = ?", true).Preload("FromCurrency").Preload("ToCurrency").Find(&exchanges).Error; err != nil {
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

func UpdateExchangeRateByID(exchangeRateId string, updateExchangeRateReq dto.UpdateExchangeRateReq) error {
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
	err := config.DB.
		Joins(`
		JOIN app_currencies fc
		ON fc.id = app_exchange_rates.from_currency_id
		AND fc.is_active = true
	`).
		Joins(`
		JOIN app_currencies tc
		ON tc.id = app_exchange_rates.to_currency_id
		AND tc.is_active = true
	`).
		Where(
			"app_exchange_rates.from_currency_id = ? AND app_exchange_rates.to_currency_id = ? AND app_exchange_rates.is_active = ?",
			fromId, toId, true,
		).
		Preload("FromCurrency").
		Preload("ToCurrency").
		First(&exchangeRate).Error

	if err != nil {
		log.Printf("Error GetExchangeRateFromDB: %v\n", err)
		return models.ExchangeRate{}
	}
	return exchangeRate
}
