package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Greeshmanth-Pulicallu/currency/config"
	"github.com/Greeshmanth-Pulicallu/currency/controller"
	"github.com/Greeshmanth-Pulicallu/currency/models"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting server")
	config.Connect()
	config.DB.AutoMigrate(&models.Currency{})
	config.DB.AutoMigrate(&models.ExchangeRate{})

	tables, _ := config.DB.Migrator().GetTables()
	fmt.Println("Tables:")
	for _, t := range tables {
		fmt.Println("-", t)
	}
	for _, t := range tables {
		columns, err := config.DB.Migrator().ColumnTypes(t)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\nTable: %s\n", t)
		for _, c := range columns {
			fmt.Printf("Column: %s, Type: %s, Nullable: \n", c.Name(), c.DatabaseTypeName())
		}
	}

	r := mux.NewRouter()
	r.HandleFunc("/currencies", controller.CreateNewCurrencyHandler).Methods("POST")
	r.HandleFunc("/currencies", controller.GetAllActiveCurrenciesHandler).Methods("GET")
	fmt.Println(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
