package service

import (
	"github.com/Greeshmanth-Pulicallu/currency/dto"
	"github.com/Greeshmanth-Pulicallu/currency/models"
	"github.com/Greeshmanth-Pulicallu/currency/repository"
)

func AddNewExchangeRateToDB(exchangeReq dto.CreateNewExchangeRateReq) error {
	return repository.AddNewExchangeRateToDB(exchangeReq)
}

func GetAllActiveExchangesFromDB() ([]models.ExchangeRate, error) {
	return repository.GetAllActiveExchangesFromDB()
}

func GetExchangeRatesByIDFromDB(id string) (models.ExchangeRate, error) {
	return repository.GetExchangeRatesByIDFromDB(id)
}

func UpdateExchangeRateByID(id string, updateCurrency dto.UpdateExchangeRateReq) error {
	return repository.UpdateExchangeRateByID(id, updateCurrency)
}

func DeleteExchangeRateByID(id string) error {
	return repository.DeleteExchangeRateByID(id)
}
