package handlers

import (
	"github.com/gin-gonic/gin"
	. "massimple.com/wallet-controller/pkg/models"
	"massimple.com/wallet-controller/pkg/service"
	. "massimple.com/wallet-controller/pkg/webapp/dtos"
	. "massimple.com/wallet-controller/pkg/webapp/utils"
	"net/http"
	"strconv"
)

// @Summary Get available Instruments
// @Description Returns a list of the available instruments uploaded by the client
// @ID Get Instruments
// @Produce  json
// @Success 200 {array} models.Instrument
// @Failure 400 {object} string "Illegal token"
// @Failure 404 {object} string "" "no such user"
// @Router /me/instruments [get]
func GetInstrumentsHandler(c *gin.Context) {
	jwtUser, exists := c.Get(IdentityKey)
	if !exists {
		Respond(c, http.StatusBadRequest, nil, nil)
		return
	}
	instruments, err := service.GetInstrumentsById(jwtUser.(*JwtUser).getId())
	if err != nil {
		Respond(c, http.StatusNotFound, nil, err)
		return
	}
	Respond(c, http.StatusOK, instruments, nil)
}

// @Summary Insert instrument
// @Description Inserts and instrument to the list of available user instruments
// @Description Return the instrument object with its id
// @ID Insert instrument
// @Produce  json
// @Param   id     path    uint     true    "Id of the user that requests the instruments"
// @Param   instrument body  dtos.InstrumentDto  true "Instrument to insert"
// @Success 200 {array} models.Instrument
// @Failure 400 {object} string "The id provided is illegal"
// @Failure 404 {object} string "id does not exist"
// @Router /me/instruments [post]
func InsertInstrumentsHandler(c *gin.Context) {
	user, exists := c.Get(IdentityKey)
	if !exists {
		Respond(c, http.StatusBadRequest, nil, nil)
		return
	}
	var instrumentDto InstrumentDto
	if err := c.BindJSON(&instrumentDto); err != nil {
		Respond(c, http.StatusBadRequest, "You must provide an complete instrument", nil)
		return
	}
	instrument, err := service.InsertInstrumentById(user.(*JwtUser).getId(), instrumentDto.Build())
	if err != nil {
		Respond(c, http.StatusNotFound, nil, err)
		return
	}
	Respond(c, http.StatusOK, instrument, nil)
}

// @Summary Delete an Instrument
// @Description Deletes one of the instruments available to the user
// @ID Delete Instruments
// @Produce  plain
// @Param   id     path    uint     true    "Id of the instrument to delete"
// @Success 200 {object} string "Card deleted"
// @Failure 400 {object} string "" "The id provided is illegal"
// @Failure 404 {object} string "" "id does not exist"
// @Router /me/instruments/{id} [delete]
func DeleteInstrumentsHandler(c *gin.Context) {
	user, exists := c.Get(IdentityKey)
	if !exists {
		Respond(c, http.StatusBadRequest, nil, nil)
		return
	}
	instrumentId, err := strconv.ParseUint(c.Param("id"),10, 64)
	if err != nil {
		Respond(c, http.StatusBadRequest, nil, err)
		return
	}
	err = service.DeleteInstrumentById(user.(*JwtUser).getId(), uint(instrumentId))
	if err, ok := err.(*UnauthorizedError); ok {
		Respond(c, http.StatusUnauthorized, nil, err)
	}
	if err != nil {
		Respond(c, http.StatusNotFound, nil, err)
		return
	}
	Respond(c, http.StatusOK, "Card deleted", nil)
}
