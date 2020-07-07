package handlers

import (
	"github.com/gin-gonic/gin"
	"massimple.com/wallet-controller/pkg/services"
	. "massimple.com/wallet-controller/pkg/webapp/utils"
	"net/http"
)

type PhoneNumber struct {
	PhoneNumber string `json:"phoneNumber" example:"+5491133071114"`
}

// @Summary Login
// @Description Return the user given the provided phone number
// @ID Get User
// @Produce  json
// @Param   phoneNumber body  handlers.PhoneNumber true "user's phone number"
// @Success 200 {array} models.Account
// @Failure 400 {object} string "The phone number provided is illegal"
// @Failure 404 {object} string "" "id does not exist"
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	var phoneNumber PhoneNumber
	if err := c.BindJSON(&phoneNumber); err != nil {
		Respond(c, http.StatusBadRequest, nil, nil)
		return
	}
	account := services.GetAccount(phoneNumber.PhoneNumber)
	Respond(c, http.StatusOK, account, nil)
}
