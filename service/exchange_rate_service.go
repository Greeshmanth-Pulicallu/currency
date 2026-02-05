package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/shopspring/decimal"

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

const ExchngeRateURL = "https://v6.exchangerate-api.com/v6/%v/latest/%v"

func FetchExchangeExchangeRatesFromExternalAPIService(baseCode string) error {
	activeCurrencies, err := repository.GetAllActiveCurrenciesFromDB()
	if err != nil {
		return err
	}
	baseCodeIsActive := false
	var baseCodeID uint

	for _, c := range activeCurrencies {
		if c.Code == baseCode {
			baseCodeIsActive = true
			baseCodeID = c.ID
			break
		}
	}

	if !baseCodeIsActive {
		return errors.New("Invalid base_code")
	}

	apiKey := os.Getenv("API_SECRET")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf(ExchngeRateURL, apiKey, baseCode),
		nil,
	)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	var result dto.ExchangeRateResponse
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)

	for _, c := range activeCurrencies {
		apiExchangeRate, ok := result.ConversionRates[c.Code]
		if !ok || c.Code == baseCode {
			continue
		}

		newExchangeRate := dto.CreateNewExchangeRateReq{}
		newExchangeRate.FromCurrencyID = baseCodeID
		newExchangeRate.ToCurrencyID = c.ID
		newExchangeRate.Rate = decimal.NewFromFloat(apiExchangeRate)

		if err := repository.AddNewExchangeRateToDB(newExchangeRate); err != nil {
			fmt.Printf("Error from service: %v\n", err)
		}
		fmt.Printf("added %v %v\n", baseCodeID, c.ID)

	}

	return nil
}
