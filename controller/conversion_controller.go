package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Greeshmanth-Pulicallu/currency/service"
)

func ConvertCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	amount := r.URL.Query().Get("amount")

	if to == "" || from == "" || amount == "" {
		http.Error(w, "Query params from to amount necessary", http.StatusBadRequest)
		return
	}

	response, err := service.ConvertCurrencyService(from, to, amount)
	if err != nil {
		http.Error(w, "invalid params", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
