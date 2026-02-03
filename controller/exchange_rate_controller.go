package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Greeshmanth-Pulicallu/currency/dto"
	"github.com/Greeshmanth-Pulicallu/currency/repository"
	"github.com/gin-gonic/gin"
)

func CreateNewExchangeRateHandler(c *gin.Context) {
	w := c.Writer
	r := c.Request

	var exchangeReq dto.CreateNewExchangeRateReq

	if err := json.NewDecoder(r.Body).Decode(&exchangeReq); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	if exchangeReq.FromCurrencyID == exchangeReq.ToCurrencyID {
		http.Error(w, "from and to id must be different", http.StatusBadRequest)
		return
	}

	if err := repository.AddNewExchangeRateToDB(exchangeReq); err != nil {
		http.Error(w, "exchange rate already exists", http.StatusBadRequest)
		fmt.Printf("Error: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetAllActiveExchangeRatesHandler(c *gin.Context) {
	w := c.Writer

	activeCurrencies, err := repository.GetAllActiveExchangesFromDB()
	if err != nil {
		http.Error(w, "Internal", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activeCurrencies)
}

func GetExchangeRatesByIDHandler(c *gin.Context) {
	w := c.Writer
	id := c.Param("id")

	currency, err := repository.GetExchangeRatesByIDFromDB(id)
	if err != nil {
		fmt.Printf("Error from GetExchangeRatesByIDHandler %v\n", err)
		http.Error(w, "Id not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currency)
}

func UpdateExchangeRatesByIDHandler(c *gin.Context) {
	w := c.Writer
	r := c.Request
	id := c.Param("id")

	var updateCurrency dto.UpdateExchangeRateReq

	if err := json.NewDecoder(r.Body).Decode(&updateCurrency); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	if err := repository.UpdateExchangeRateByID(id, updateCurrency); err != nil {
		fmt.Println(err)
		http.Error(w, "Id does not exist", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
}

func DeleteExchangeRatesByIDHandler(c *gin.Context) {
	w := c.Writer
	id := c.Param("id")

	if err := repository.DeleteExchangeRateByID(id); err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
}
