package repository

import (
	"errors"
	"log"

	"github.com/Greeshmanth-Pulicallu/currency/config"
	"github.com/Greeshmanth-Pulicallu/currency/dto"
	"github.com/Greeshmanth-Pulicallu/currency/models"
)

func AddNewCurrencyToDB(currencyFromReq dto.CreateNewCurrencyReq) error {
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

func UpdateCurrencyByID(currencyId string, updateReq dto.UpdateCurrencyReq) error {
	result := config.DB.
		Model(&models.Currency{}).
		Where("id = ?", currencyId).
		Updates(updateReq)

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

func GetCurrencyFromDB(code string) ([]models.Currency, error) {
	var result []models.Currency
	if err := config.DB.Where("code = ?", code).Find(&result).Error; err != nil {
		log.Printf("Error GetCurrencyFromDB: %v\n", err)
		return []models.Currency{}, err
	}
	return result, nil
}
