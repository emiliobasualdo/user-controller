package webapp

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	. "massimple.com/wallet-controller/internal/webapp/handlers"
)


func Serve() {
	if viper.GetBool("server.verbose") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	// Creates a gin router with default middleware:
	router := gin.Default()
	authMiddleware, err := AuthMiddleware()
	if err != nil {
		panic(err)
	}
	health := router.Group("/")
	{
		health.GET("", StatusHandler)
	}
	auth := router.Group("/auth")
	auth.POST("/sms-code", SendSmsHandler)
	auth.POST("/login", authMiddleware.LoginHandler)
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	me := router.Group("/me", authMiddleware.MiddlewareFunc(), AuthMiddlewareWrapper())
	{
		me.GET("", MeHandler)
		me.POST("", EditMeHandler)
		me.GET("/home", HomeHandler)
		instruments := me.Group("/instruments")
		{
			instruments.GET("", GetInstrumentsHandler)
			instruments.POST("", InsertInstrumentsHandler)
		}
		transactions := me.Group("/transactions")
		{
			transactions.GET("", TransactionHistoryHandler)
			transactions.POST("", NewTransactionHandler)
		}
	}
	// swagger docs
	router.GET("/api-doc/*any", SwaggerHandler)

	port := viper.GetString("server.port")
	if err := router.Run(":"+port); err != nil {
		panic(err)
	}
}
