package api

import (

	"payment-system-one/internal/models"
	"payment-system-one/internal/util"
	"github.com/gin-gonic/gin"

)

// Create a user
func (u *HTTPHandler) CreateUser(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBind(&user); err != nil {
		util.Response(c, "invalid request 123", 400, "bad request body", nil)
		return
	}
	//validate user email
	if  !util.IsValidEmail(user.Email) {
		util.Response(c, "Invalid email", 400,  "Bad request body", nil)
		return
	}
	//check if user already exists
	_, err := u.Repository.FindUserByEmail(user.Email)
	if err == nil {
		
		util.Response(c, "User already exists", 400, "Bad request body", nil)
		return
	}

	hashPass, err := util.HashPassword(user.Password)
	if err != nil {
		util.Response(c, "could not hash password", 500, "internal server error", nil)
		return
	}

	user.Password = hashPass

	//Generate account number
	acctNo, err := util.GenerateAccountNumber()
	if err != nil {
		util.Response(c, "could not generate account number", 500, "internal server error", nil)
		return
	}

	user.AccountNo = acctNo

	//set available balance to zero
	user.AvailableBalance = 0.0

	//Persist the user on the database
	err = u.Repository.CreateUser(user)
	if err != nil {
		util.Response(c, "user not created", 400, err.Error(), nil)
		return
	}
	util.Response(c, "user created", 200, "success", nil)
}