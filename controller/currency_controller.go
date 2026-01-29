package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Greeshmanth-Pulicallu/currency/config"
	"github.com/Greeshmanth-Pulicallu/currency/repository"
)

func CreateNewCurrencyHandler(w http.ResponseWriter, r *http.Request) {
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
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(currency)

}
