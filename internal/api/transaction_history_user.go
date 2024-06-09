package api

import (
	
	"payment-system-five/internal/util"
	"github.com/gin-gonic/gin"
)

func (u *HTTPHandler) ViewUserTransactionsHistory(c *gin.Context) {
	//var transaction *[]models.Transaction
	/* var accountNo *models.UserAccountNo
	if err := c.ShouldBind(&accountNo); err != nil {
		util.Response(c, "invalid request", 400, "bad request body", nil)
		return
	} */

	//Get user from context
	user, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "User not logged in", 500, "user not found", nil)
		return
	}

	// Allow user to see there transaction
	usertransaction, err := u.Repository.GetTransactionByAccountNumber(user.AccountNo)
	if err != nil {
		util.Response(c, "Transaction not found", 400, "Bad request body", nil)
		return 
	}
		util.Response(c, "Transaction found", 200, usertransaction, nil)
		return
		


	// persist the data into the db
	/* err = u.Repository.ViewUserTransactionsHistory(transaction)
	if err != nil {
		util.Response(c, "Transaction history not find", 500, "Demand not successful", nil)
		return

	}
	util.Response(c, "Transaction history found", 200, "Demand succesful", nil) */




	



}