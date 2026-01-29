package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Greeshmanth-Pulicallu/currency/config"
	"github.com/Greeshmanth-Pulicallu/currency/repository"
)

func CreateNewExchangeRateHandler(w http.ResponseWriter, r *http.Request) {
	var exchangeReq config.CreateNewExchangeRateReq

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
		http.Error(w, "Internal", http.StatusInternalServerError)
		fmt.Printf("Error: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetAllActiveExchangeRatesHandler(w http.ResponseWriter, r *http.Request) {
	activeCurrencies, err := repository.GetAllActiveExchangesFromDB()
	if err != nil {
		http.Error(w, "Internal", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activeCurrencies)
}

func GetExchangeRatesByIDHandler(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Println("HIT")
	currency, err := repository.GetExchangeRatesByIDFromDB(id)
	if err != nil {
		fmt.Printf("Error from GetCurrencyByIDHandler %v\n", err)
		http.Error(w, "Id not found", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currency)

}
