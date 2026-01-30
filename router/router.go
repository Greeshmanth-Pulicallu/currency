package router

import (
	"github.com/Greeshmanth-Pulicallu/currency/controller"
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func init() {
	R = gin.Default()

	// currencies
	R.POST("/currencies", controller.CreateNewCurrencyHandler)
	R.GET("/currencies", controller.GetAllActiveCurrenciesHandler)
	R.GET("/currencies/:id", controller.GetCurrencyByIDHandler)
	R.PUT("/currencies/:id", controller.UpdateCurrencyByIDHandler)
	R.DELETE("/currencies/:id", controller.DeleteCurrencyByIDHandler)

	// exchange rates
	R.POST("/exchange-rates", controller.CreateNewExchangeRateHandler)
	R.GET("/exchange-rates", controller.GetAllActiveExchangeRatesHandler)
	R.GET("/exchange-rates/:id", controller.GetExchangeRatesByIDHandler)
	R.PUT("/exchange-rates/:id", controller.UpdateExchangeRatesByIDHandler)
	R.DELETE("/exchange-rates/:id", controller.DeleteExchangeRatesByIDHandler)

	// conversion
	R.GET("/convert", controller.ConvertCurrencyHandler)
}
