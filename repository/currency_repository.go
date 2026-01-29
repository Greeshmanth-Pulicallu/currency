package repository

import (
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
