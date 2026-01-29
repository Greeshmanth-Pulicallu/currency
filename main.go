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
	r.HandleFunc("currencies/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		controller.GetCurrencyByIDHandler(w, r, id)
	}).Methods("GET")
	r.HandleFunc("/currencies/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		controller.UpdateCurrencyByIDHandler(w, r, id)
	}).Methods("PUT")
	r.HandleFunc("/exchange-rates/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		controller.DeleteCurrencyByIDHandler(w, r, id)
	}).Methods("DELETE")

	r.HandleFunc("/exchange-rates", controller.CreateNewExchangeRateHandler).Methods("POST")
	r.HandleFunc("/exchange-rates", controller.GetAllActiveExchangeRatesHandler).Methods("GET")
	r.HandleFunc("/exchange-rates/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		controller.GetExchangeRatesByIDHandler(w, r, id)
	}).Methods("GET")
	r.HandleFunc("/exchange-rates/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		controller.UpdateExchangeRatesByIDHandler(w, r, id)
	}).Methods("PUT")
	r.HandleFunc("/exchange-rates/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		controller.DeleteExchangeRatesByIDHandler(w, r, id)
	}).Methods("DELETE")

	r.HandleFunc("/convert", controller.ConvertCurrencyHandler).Methods("GET")

	fmt.Println(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
