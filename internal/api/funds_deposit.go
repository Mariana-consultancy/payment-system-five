package api

import (
	"payment-system-five/internal/models"
	"payment-system-five/internal/util"

	"github.com/gin-gonic/gin"
)

func (u *HTTPHandler) AddFunds(c *gin.Context) {
	var funds *models.AddFunds
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

	user.AvailableBalance += funds.Amount

	// persist the data into the db
	err = u.Repository.UpdateUser(user)
	if err != nil {
		util.Response(c, "could not add funds", 500, "could not add funds", nil)
		return
	}
	util.Response(c, "funds were added successfully", 200, "Transfer successful", nil)

}
