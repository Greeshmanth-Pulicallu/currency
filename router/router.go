package router

import (
	"github.com/Greeshmanth-Pulicallu/currency/controller"
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func init() {
	R = gin.Default()

	protected := R.Group("/")
	protected.Use(controller.JWTAuthMiddleware())

	// currencies
	protected.POST("/currencies", controller.CreateNewCurrencyHandler)
	protected.GET("/currencies", controller.GetAllActiveCurrenciesHandler)
	protected.GET("/currencies/:id", controller.GetCurrencyByIDHandler)
	protected.PUT("/currencies/:id", controller.UpdateCurrencyByIDHandler)
	protected.DELETE("/currencies/:id", controller.DeleteCurrencyByIDHandler)

	// exchange rates
	protected.POST("/exchange-rates", controller.CreateNewExchangeRateHandler)
	protected.GET("/exchange-rates", controller.GetAllActiveExchangeRatesHandler)
	protected.GET("/exchange-rates/:id", controller.GetExchangeRatesByIDHandler)
	protected.PUT("/exchange-rates/:id", controller.UpdateExchangeRatesByIDHandler)
	protected.DELETE("/exchange-rates/:id", controller.DeleteExchangeRatesByIDHandler)
	protected.GET("/exchange-rates/fetch-latest-rates/:base_code", controller.FetchExchangeRatesFromExternalAPI)

	// conversion
	protected.GET("/convert", controller.ConvertCurrencyHandler)

	// auth
	R.POST("/auth/register", controller.UserRegisterHandler)
	R.POST("/auth/login", controller.UserLoginHandler)
}
