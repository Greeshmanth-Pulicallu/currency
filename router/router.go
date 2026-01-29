package router

import (
	"net/http"

	"github.com/Greeshmanth-Pulicallu/currency/controller"
	"github.com/gorilla/mux"
)

var R *mux.Router

func init() {
	R = mux.NewRouter()
	R.HandleFunc("/currencies", controller.CreateNewCurrencyHandler).Methods("POST")
	R.HandleFunc("/currencies", controller.GetAllActiveCurrenciesHandler).Methods("GET")
	R.HandleFunc("currencies/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		controller.GetCurrencyByIDHandler(w, r, id)
	}).Methods("GET")
	R.HandleFunc("/currencies/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		controller.UpdateCurrencyByIDHandler(w, r, id)
	}).Methods("PUT")
	R.HandleFunc("/exchange-rates/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		controller.DeleteCurrencyByIDHandler(w, r, id)
	}).Methods("DELETE")

	R.HandleFunc("/exchange-rates", controller.CreateNewExchangeRateHandler).Methods("POST")
	R.HandleFunc("/exchange-rates", controller.GetAllActiveExchangeRatesHandler).Methods("GET")
	R.HandleFunc("/exchange-rates/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		controller.GetExchangeRatesByIDHandler(w, r, id)
	}).Methods("GET")
	R.HandleFunc("/exchange-rates/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		controller.UpdateExchangeRatesByIDHandler(w, r, id)
	}).Methods("PUT")
	R.HandleFunc("/exchange-rates/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		controller.DeleteExchangeRatesByIDHandler(w, r, id)
	}).Methods("DELETE")

	R.HandleFunc("/convert", controller.ConvertCurrencyHandler).Methods("GET")
}
