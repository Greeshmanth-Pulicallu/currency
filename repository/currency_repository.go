package repository

import (
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
