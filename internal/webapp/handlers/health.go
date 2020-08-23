package handlers

import (
	"github.com/gin-gonic/gin"
	. "massimple.com/wallet-controller/internal/webapp/utils"
	"net/http"
)

// @Summary Health check
// @Description Returns a small body to ensure the server is up and running
// @query Get User
// @Produce  json
// @Success 200 {object} handlers.health
// @Failure 400 {object}
// @Failure 500 {object}
// @Router /auth/login [post]

type health struct {
	Status	string `json:"status" example:"ok"`
	Version string `json:"version" example:"v1"`
}

func StatusHandler(c *gin.Context) {
	Respond(c, http.StatusOK, health{
		Status:  "ok",
		Version: "v1",
	}, nil)
}