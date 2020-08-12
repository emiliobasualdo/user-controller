package handlers

import (
	"github.com/gin-gonic/gin"
	. "massimple.com/wallet-controller/internal/dtos"
	"massimple.com/wallet-controller/internal/service"
	. "massimple.com/wallet-controller/internal/webapp/utils"
	"net/http"
)

// @Summary Get account details
// @Description Returns an account by its id
// @query Get account
// @Produce  json
// @Success 200 {object} models.Account
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
	Respond(c, http.StatusOK, AccountDtoFromAccount(acc), nil)
}

// @Summary Edit account information
// @Description Replaces all the account information with the information pased
// @query Edit account
// @Produce  json
// @Param   account body  dtos.AccountDto  true "Fields to edit"
// @Success 200 {object}
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "" "no such user"
// @Router /me [post]
func EditMeHandler(c *gin.Context)  {
	user, _ := c.Get(IdentityKey)
	var accountDto AccountDto
	if err := c.BindJSON(&accountDto); err != nil {
		Respond(c, http.StatusBadRequest, "You must provide an complete Account Dto", nil)
		return
	}
	err := service.EditAccount(user.(*JwtUser).getId(), accountDto)
	if err != nil {
		Respond(c, http.StatusNotFound, nil, err)
		return
	}
	Respond(c, http.StatusOK, nil, nil)
}
