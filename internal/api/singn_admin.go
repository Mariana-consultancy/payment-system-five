package api

import (
	"github.com/gin-gonic/gin"
	"payment-system-five/internal/models"
	"payment-system-five/internal/util"
)

// Create an admin

func (u *HTTPHandler) CreateAdmin(c *gin.Context) {
	var admin *models.Admin
	if err := c.ShouldBind(&admin); err != nil {
		util.Response(c, "invalid request", 400, "bad request body", nil)
		return
	}
	//validate admin email
	if !util.IsValidEmail(admin.Email) {
		util.Response(c, "Invalid email", 400, "Bad request body", nil)
		return
	}
	//check if admin already exists
	_, err := u.Repository.FindAdminByEmail(admin.Email)
	if err == nil {

		util.Response(c, "Admin already exists", 400, "Bad request body", nil)
		return
	}

	hashPass, err := util.HashPassword(admin.Password)
	if err != nil {
		util.Response(c, "could not hash password", 500, "internal server error", nil)
		return
	}

	admin.Password = hashPass

	err = u.Repository.CreateAdmin(admin)
	if err != nil {
		util.Response(c, "admin not created", 400, err.Error(), nil)
		return
	}
	util.Response(c, "admin created", 200, "success", nil)
}
