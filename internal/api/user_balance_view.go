package api

import (
	"payment-system-five/internal/models"
	"payment-system-five/internal/util"

	"github.com/gin-gonic/gin"
)

func (u *HTTPHandler) ViewUserBalance(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBind(&user); err != nil {
		util.Response(c, "invalid request", 400, "bad request body", nil)
		return
	}

	//Get user from context
	user, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "User not logged in", 500, "user not found", nil)
		return
	}

	//check if the account number exist
	user, err = u.Repository.FindUserByAccountNumber(user.AccountNo)
	if err != nil {
		util.Response(c, "User not found", 400, "Bad request body", nil)
		return
	}

	//Validate user
	if user.AvailableBalance != user.AvailableBalance {
		util.Response(c, "Invalid user balance", 500, "Bad request body", nil)
		return
	}

	// persist the data into the db
	/* err = u.Repository.ViewUserBalance(user)
	if err != nil {
		util.Response(c, "User balance not found", 500, "User balance unavailable", nil)
		return

	} */
	util.Response(c, "User balance found", 200, "User balance available", nil)

}
