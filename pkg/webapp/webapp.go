package webapp

import (
	"github.com/apsdehal/go-logger"
	"github.com/gin-gonic/gin"
	. "massimple.com/wallet-controller/pkg/webapp/handlers"
)

const PORT = "5000"

var log *logger.Logger

func Serve(_log *logger.Logger) {
	log = _log
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
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
		me.GET("/", MeHandler)
		me.POST("/", EditMeHandler)
		instruments := me.Group("/instruments")
		{
			instruments.GET("/", GetInstrumentsHandler)
			instruments.POST("/", InsertInstrumentsHandler)
			instruments.DELETE("/:id", DeleteInstrumentsHandler)
		}
		transactions := me.Group("/transactions")
		{
			transactions.GET("/", TransactionHistoryHandler)
			transactions.POST("/", NewTransactionHandler)
		}
	}
	// swagger docs
	router.GET("/api-doc/*any", SwaggerHandler)

	if err := router.Run(":"+PORT); err != nil {
		panic(err)
	}
	log.NoticeF("Server started on port %s", PORT)
}
