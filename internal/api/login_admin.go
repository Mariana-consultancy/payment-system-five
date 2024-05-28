package api

import (
	
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"payment-system-one/internal/middleware"
	"payment-system-one/internal/models"
	"payment-system-one/internal/util"
	"time"
)






// Create a login system for an Admin
//Login

func (u *HTTPHandler) LoginAdmin(c *gin.Context) {
	var adminloginRequest *models.AdminLoginRequest
	if err := c.ShouldBind(&adminloginRequest); err != nil {
		util.Response(c, "invalid request", 400, "bad request body", nil)
		return
	}
	if adminloginRequest.Email == "" || adminloginRequest.Password == "" {
		util.Response(c, "Please enter your email or password", 400, "bad request body", nil)
		return
	}

	admin, err := u.Repository.FindAdminByEmail(adminloginRequest.Email)
	if err != nil {
		util.Response(c, "admin does not exist", 404, "admin not found", nil)
		return
	}
	if admin.LoginCounter >= 3 {
		admin.IsLocked = true
		admin.UpdatedAt = time.Now()
		err = u.Repository.UpdateAdmin(admin)
		if err != nil {
			return
		}
		util.Response(c, "Your account has been lock after 3 failed attempt, contact customer care for assistance", 200, "success", nil)
		return
	}

	hashPass, err := util.HashPassword(admin.Password)
	if err != nil {
		util.Response(c, "could not hash password", 500, "internal server error", nil)
		return
	}

	admin.Password = hashPass

	if admin.Password != adminloginRequest.Password {
		admin.LoginCounter++
		err := u.Repository.UpdateAdmin(admin)
		if err != nil {
			util.Response(c, "internal server error", 500, "admin not found", nil)
			return
		}
		util.Response(c, "password mismatch", 404, "admin not found", nil)
		return
	}

	//Generate token
	accessClaims, refreshClaims := middleware.GenerateClaims(admin.Email)

	secret := os.Getenv("JWT_SECRET")

	accessToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)
	if err != nil {
		util.Response(c, "error generating access token", 500, "error generating access token", nil)
		return
	}
	refreshToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, refreshClaims, &secret)
	if err != nil {
		util.Response(c, "error generating refresh token", 500, "error generating refresh token", nil)
		return
	}
	c.Header("access_token", *accessToken)
	c.Header("refresh_token", *refreshToken)

	util.Response(c, "login successful", http.StatusOK, gin.H{
		"admin":          admin,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil)
}
