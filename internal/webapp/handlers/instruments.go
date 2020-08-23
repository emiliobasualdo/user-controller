package handlers

import (
	"github.com/gin-gonic/gin"
	"massimple.com/wallet-controller/internal/service"
	. "massimple.com/wallet-controller/internal/dtos"
	. "massimple.com/wallet-controller/internal/webapp/utils"
	"net/http"
)

// @Summary Get available Instruments
// @Description Returns a list of the available instruments uploaded by the client
// @query Get Instruments
// @Produce  json
// @Success 200 {array} models.Instrument
// @Failure 401 "Unauthorized"
// @Router /me/instruments [get]
func GetInstrumentsHandler(c *gin.Context) {
	jwtUser, _ := c.Get(IdentityKey)
	instruments, err := service.GetEnabledInstrumentsByAccountId(jwtUser.(*JwtUser).getId())
	if err != nil {
		Respond(c, http.StatusNotFound, nil, err)
		return
	}
	Respond(c, http.StatusOK, instruments, nil)
}

// @Summary Insert instrument
// @Description Inserts and instrument to the list of available user instruments
// @Description Return the instrument object with its id
// @query Insert instrument
// @Produce  json
// @Param   id     path    uint     true    "ID of the user that requests the instruments"
// @Param   instrument body  dtos.InstrumentDto  true "Instrument to insert"
// @Success 200
// @Failure 401 "Unauthorized"
// @Router /me/instruments [post]
func InsertInstrumentsHandler(c *gin.Context) {
	jwtUser, _ := c.Get(IdentityKey)
	var instrumentDto InstrumentDto
	if err := c.BindJSON(&instrumentDto); err != nil {
		Respond(c, http.StatusBadRequest, "You must provide an complete instrument", nil)
		return
	}
	err := service.InsertInstrumentById(jwtUser.(*JwtUser).getId(), instrumentDto.Build())
	if err != nil {
		Respond(c, http.StatusNotFound, nil, err)
		return
	}
	Respond(c, http.StatusOK, nil, nil)
}

