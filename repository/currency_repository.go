package repository

import (
	"errors"
	"log"

	"github.com/Greeshmanth-Pulicallu/currency/config"
	"github.com/Greeshmanth-Pulicallu/currency/models"
)

func AddNewCurrencyToDB(currencyFromReq config.CreateNewCurrencyReq) error {
	var newCurrency models.Currency

	newCurrency.Code = currencyFromReq.Code
	newCurrency.Name = currencyFromReq.Name
	newCurrency.Symbol = currencyFromReq.Symbol
	newCurrency.IsActive = true

	if err := config.DB.Create(&newCurrency).Error; err != nil {
		return err
	}

	return nil
}

func GetAllActiveCurrenciesFromDB() ([]models.Currency, error) {
	var currencies []models.Currency
	if err := config.DB.Where("is_active = ?", true).Find(&currencies).Error; err != nil {
		log.Printf("Error GetAllActiveCurrenciesFromDB: %v\n", err)
		return []models.Currency{}, err
	}

	return currencies, nil
}

func GetCurrencyByIDFromDB(currencyId string) (models.Currency, error) {
	var currency models.Currency
	if err := config.DB.Where("id = ?", currencyId).First(&currency).Error; err != nil {
		return models.Currency{}, nil
	}
	return currency, nil
}

func UpdateCurrencyByID(currencyId string, updateReq config.UpdateCurrencyReq) error {
	result := config.DB.
		Model(&models.Currency{}).
		Where("id = ?", currencyId).
		Updates(map[string]any{
			"name":      updateReq.Name,
			"symbol":    updateReq.Symbol,
			"is_active": updateReq.IsActive,
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Id does not exist")
	}
	return nil
}

func DeleteCurrencyByID(currencyId string) error {
	result := config.DB.
		Model(&models.Currency{}).
		Where("id = ?", currencyId).
		Updates(map[string]any{
			"is_active": false,
		})

	return result.Error
}
