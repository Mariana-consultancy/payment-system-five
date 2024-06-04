package api

import (
	"payment-system-one/internal/util"
	"payment-system-one/internal/models"
	"github.com/gin-gonic/gin"
)

// declare request body

	//bind JSON data to struct

	func (u *HTTPHandler) DepositFunds(c *gin.Context) {
		var deposit *models.Deposit
		if err := c.ShouldBind(&deposit); err != nil {
			util.Response(c, "invalid request", 400, "bad request body", nil)
			return
		}

		//Get user from context

user, err := u.GetUserFromContext(c)

if err != nil {
	util.Response(c, "User not logged in", 500, "user not found", nil)
}

//Validate the amount

addfunds := 
if addfunds <= 0 {
	util.Response(c, "Invalid amount", 500, "Bad request body", nil)
	return
}	



//check if the account number exist 

recipient, err := u.Repository.FindUserDepositByAccountNumber(addfunds.RecipiencACC)
	if err == nil {
		
		util.Response(c, "User not found", 400, "Bad request body", nil)
		return
	}


//Persist the data on the database
err = u.Repository.DepositFunds(user, recipient, addfunds.Amount )
if err != nil {
	util.Response(c, "funds not deposited", 500, "Deposit of the funds not successful", nil)
	return
}
util.Response(c, "funds deposited into the account", 200, "Add funds successful", nil)

}