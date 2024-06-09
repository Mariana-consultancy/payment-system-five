package api

import (
	"payment-system-five/internal/models"
	"payment-system-five/internal/util"

	"github.com/gin-gonic/gin"
)

func (u *HTTPHandler) AdminViewAllTransactionsHistory(c *gin.Context) {
	var transaction *models.Transaction
	if err := c.ShouldBind(&transaction); err != nil {
		util.Response(c, "invalid request", 400, "bad request body", nil)
		return
	}

	//Get user from context
	_, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "User not logged in", 500, "user not found", nil)
		return
	}

	//Allow admin to see transaction of all user
	if transaction.PayerAccount != transaction.PayerAccount {
		util.Response(c, "User transaction not found", 500, "Bad body request", nil)
		return
	}
	if transaction.RecipientAccount != transaction.RecipientAccount {
		util.Response(c, "User transaction not found", 500, "Bad body request", nil)
		return
	}

	// persist the data into the db
	/* err = u.Repository.AdminViewAllTransactionsHistory(transaction)
	if err != nil {
		util.Response(c, "transaction not found", 500, "transaction history unavailable", nil)
		return

	} */
	util.Response(c, "transaction found", 200, "transaction history available", nil)

}