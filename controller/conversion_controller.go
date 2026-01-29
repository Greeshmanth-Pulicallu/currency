package controller

import (
	"encoding/json"
	"net/http"

	"github.com/shopspring/decimal"

	"github.com/Greeshmanth-Pulicallu/currency/config"
	"github.com/Greeshmanth-Pulicallu/currency/repository"
)

func ConvertCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	amount := r.URL.Query().Get("amount")

	if to == "" || from == "" || amount == "" {
		http.Error(w, "Query params from to amount necessary", http.StatusBadRequest)
		return
	}

	fromQuery, err := repository.GetCurrencyFromDB(from)
	if err != nil {
		http.Error(w, "From not found", http.StatusBadRequest)
		return
	}

	toQuery, err := repository.GetCurrencyFromDB(to)
	if err != nil {
		http.Error(w, "to not found", http.StatusBadRequest)
		return
	}

	if len(fromQuery) == 0 || len(toQuery) == 0 {
		http.Error(w, "Query params not found in db", http.StatusBadRequest)
		return
	}

	exchangeRate := repository.GetExchangeRatesForPairFromDB(fromQuery[0].ID, toQuery[0].ID)

	respAmount, err := decimal.NewFromString(amount)
	if err != nil {
		http.Error(w, "Internal", http.StatusInternalServerError)
		return
	}

	var response config.CurrencyConversionRes
	response.From = from
	response.To = to
	response.Amount = respAmount
	response.ExchangeRate = exchangeRate.Rate
	response.ConvertedAmount = respAmount.Mul(exchangeRate.Rate)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
