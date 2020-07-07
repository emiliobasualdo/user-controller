package handlers

import (
	"github.com/gin-gonic/gin"
	. "massimple.com/wallet-controller/pkg/webapp/utils"
	"net/http"
)

func ExecutionHandler(c *gin.Context) {
	resp := "No implementado"
	Respond(c, http.StatusNotImplemented, resp, nil)
}
