package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Greeshmanth-Pulicallu/currency/config"
	"github.com/Greeshmanth-Pulicallu/currency/repository"
	"github.com/gin-gonic/gin"
)

func CreateNewCurrencyHandler(c *gin.Context) {
	w := c.Writer
	r := c.Request
	var currency config.CreateNewCurrencyReq

	if err := json.NewDecoder(r.Body).Decode(&currency); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// ensure code is uppercase
	if currency.Code != strings.ToUpper(currency.Code) {
		http.Error(w, "currency code must be uppercase", http.StatusBadRequest)
		return
	}

	if err := repository.AddNewCurrencyToDB(currency); err != nil {
		http.Error(w, "Unable to add to db", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(currency)

}

func GetAllActiveCurrenciesHandler(c *gin.Context) {
	w := c.Writer
	// r := c.Request
	activeCurrencies, err := repository.GetAllActiveCurrenciesFromDB()
	if err != nil {
		http.Error(w, "Internal", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activeCurrencies)
}

func GetCurrencyByIDHandler(c *gin.Context) {
	w := c.Writer
	id := c.Param("id")

	currency, err := repository.GetCurrencyByIDFromDB(id)
	if err != nil {
		fmt.Printf("Error from GetCurrencyByIDHandler %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currency)

}

func UpdateCurrencyByIDHandler(c *gin.Context) {
	w := c.Writer
	r := c.Request
	id := c.Param("id")

	var updateCurrency config.UpdateCurrencyReq

	if err := json.NewDecoder(r.Body).Decode(&updateCurrency); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	if err := repository.UpdateCurrencyByID(id, updateCurrency); err != nil {
		fmt.Println(err)
		http.Error(w, "Id does not exist", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
}

func DeleteCurrencyByIDHandler(c *gin.Context) {
	w := c.Writer
	id := c.Param("id")

	if err := repository.DeleteCurrencyByID(id); err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
}
