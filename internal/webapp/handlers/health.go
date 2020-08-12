package handlers

import (
	"github.com/gin-gonic/gin"
	. "massimple.com/wallet-controller/internal/webapp/utils"
	"net/http"
)

func StatusHandler(c *gin.Context) {
	Respond(c, http.StatusOK, gin.H{
		"status":  "ok",
		"version": "v1",
		"host": c.Request.Host,
	}, nil)
}