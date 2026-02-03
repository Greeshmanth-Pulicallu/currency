package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Greeshmanth-Pulicallu/currency/service"
	"github.com/gin-gonic/gin"
)

func ConvertCurrencyHandler(c *gin.Context) {
	w := c.Writer

	from := c.Query("from")
	to := c.Query("to")
	amount := c.Query("amount")

	if from == "" || to == "" || amount == "" {
		http.Error(w, "Query params from to amount necessary", http.StatusBadRequest)
		return
	}

	response, err := service.ConvertCurrencyService(from, to, amount)
	if err != nil {
		http.Error(w, "valid ids and non zero amount is required", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
