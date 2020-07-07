package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"massimple.com/wallet-controller/docs"
)

func SwaggerHandler(c *gin.Context) {
	docs.SwaggerInfo.Host = c.Request.Host
	ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
}
