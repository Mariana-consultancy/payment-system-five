package api

import (
	"github.com/gin-gonic/gin"
	"payment-system-one/internal/models"
	"payment-system-one/internal/util"
)

// declare request body

//bind JSON data to struct

func (u *HTTPHandler) TransferFunds(c *gin.Context) {
	var funds *models.TransferFunds
	if err := c.ShouldBind(&funds); err != nil {
		util.Response(c, "invalid request", 400, "bad request body", nil)
		return
	}

	//Get user from context

	user, err := u.GetUserFromContext(c)

	if err != nil {
		util.Response(c, "User not logged in", 500, "user not found", nil)
		return
	}

	//Validate the amount
	if funds.Amount <= 0 {
		util.Response(c, "Invalid amount", 500, "Bad request body", nil)
		return
	}

	//check if the account number exist
	recipient, err := u.Repository.FindUserByAccountNumber(funds.RecipienctACC)
	if err != nil {
		util.Response(c, "User not found", 400, "Bad request body", nil)
		return
	}

	//check if amount being transferred is less than the user's current balance
	if funds.Amount >= user.AvailableBalance {
		util.Response(c, "Balance issuficient funds", 400, "bad Request", nil)
		return
	}

	// persist the data into the db
	err = u.Repository.TransferFunds(user, recipient, funds.Amount)
	if err != nil {
		util.Response(c, "Transfer not possible", 500, "transfer not succesful", nil)
		return

	}
	util.Response(c, "Transfer was done successfully", 200, "Transfer successful", nil)

}
