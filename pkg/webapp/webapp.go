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
	health := router.Group("/")
	{
		health.GET("", StatusHandler)
	}
	auth := router.Group("/login")
	{
		auth.POST("/", LoginHandler)
	}
	me := router.Group("/me")
	{
		instruments := me.Group("/instruments")
		{
			instruments.GET("/:id", GetInstrumentsHandler)
			instruments.POST("/:id", InsertInstrumentsHandler)
			instruments.DELETE("/:id", DeleteInstrumentsHandler)
		}
		transactions := me.Group("/transactions")
		{
			transactions.GET("/", ExecutionHandler)
			transactions.POST("/", ExecutionHandler)
		}
	}
	// swagger docs
	router.GET("/api-doc/*any", SwaggerHandler)

	if err := router.Run(":"+PORT); err != nil {
		panic(err)
	}
	log.NoticeF("Server started on port %s", PORT)
}