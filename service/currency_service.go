package service

import (
	"github.com/Greeshmanth-Pulicallu/currency/dto"
	"github.com/Greeshmanth-Pulicallu/currency/models"
	"github.com/Greeshmanth-Pulicallu/currency/repository"
)

func CreateNewCurrencyService(currency dto.CreateNewCurrencyReq) error {
	return repository.AddNewCurrencyToDB(currency)
}

func GetAllActiveCurrenciesFromDB() ([]models.Currency, error) {
	return repository.GetAllActiveCurrenciesFromDB()
}

func GetCurrencyByIDFromDB(id string) (models.Currency, error) {
	return repository.GetCurrencyByIDFromDB(id)
}

func UpdateCurrencyByID(id string, updateCurrency dto.UpdateCurrencyReq) error {
	return repository.UpdateCurrencyByID(id, updateCurrency)
}

func DeleteCurrencyByID(id string) error {
	return repository.DeleteCurrencyByID(id)
}
