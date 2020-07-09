package handlers

import (
	"github.com/gin-gonic/gin"
	"massimple.com/wallet-controller/pkg/service"
	. "massimple.com/wallet-controller/pkg/webapp/utils"
	"net/http"
)

// @Summary Get available Instruments
// @Description Returns a list of the available instruments uploaded by the client
// @ID Get Instruments
// @Produce  json
// @Success 200 {array} models.Account
// @Failure 400 {object} string "Illegal token"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "" "no such user"
// @Router /me [get]
func MeHandler(c *gin.Context)  {
	user, _ := c.Get(IdentityKey)
	acc, err := service.GetAccountById(user.(*JwtUser).getId())
	if err != nil {
		Respond(c, http.StatusNotFound, nil, err)
		return
	}
	Respond(c, http.StatusOK, acc, nil)
}
