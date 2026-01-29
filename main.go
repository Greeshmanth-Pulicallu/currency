package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Greeshmanth-Pulicallu/currency/config"
	"github.com/Greeshmanth-Pulicallu/currency/models"
	"github.com/Greeshmanth-Pulicallu/currency/router"
)

func main() {
	fmt.Println("Starting server")
	config.Connect()
	config.DB.AutoMigrate(&models.Currency{})
	config.DB.AutoMigrate(&models.ExchangeRate{})

	http.Handle("/", router.R)
	log.Fatal(http.ListenAndServe(":8080", router.R))
}
