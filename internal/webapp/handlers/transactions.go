package handlers

import (
	"github.com/gin-gonic/gin"
	. "massimple.com/wallet-controller/internal/dtos"
	"massimple.com/wallet-controller/internal/service"
	. "massimple.com/wallet-controller/internal/webapp/utils"
	"net/http"
)

// @Summary Get transaction history
// @Description Returns all time transactions history
// @query Get transactions history
// @Produce  json
// @Success 200 {array} models.Transaction
// @Failure 404 {object} string "No such user"
// @Router /me/transactions [get]
func TransactionHistoryHandler(c *gin.Context) {
	// todo paging
	user, _ := c.Get(IdentityKey)
	trans, err := service.GetTransactions(user.(*JwtUser).getId())
	if err != nil {
		Respond(c, http.StatusNotFound, nil, err)
		return
	}
	Respond(c, http.StatusOK, trans, nil)
}

// @Summary execute transaction
// @Description Executes as transaction and returns it's full details
// @query execute transaction
// @Produce  json
// @Param   transaction body  dtos.TransactionDto  true "Transaction to execute to insert"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} string "The transaction provided is illegal"
// @Router /me/transactions [post]
func NewTransactionHandler(c *gin.Context) {
	user, _ := c.Get(IdentityKey)
	var transactionDto TransactionDto
	if err := c.BindJSON(&transactionDto); err != nil {
		Respond(c, http.StatusBadRequest, "You must provide an complete transaction", nil)
		return
	}
	err := service.ExecuteTransaction(user.(*JwtUser).getId(), transactionDto.Build())
	if err != nil {
		Respond(c, http.StatusBadRequest, nil, err)
		return
	}
	Respond(c, http.StatusCreated, nil, nil)
}
