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





//Create a Login System for user
// login

func (u *HTTPHandler) NewLoginUser(c *gin.Context) {
	var loginRequest *models.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		util.Response(c, "invalid request", 400, "bad request body", nil)
		return
	}
	
	if loginRequest.Email == "" || loginRequest.Password == "" {
		util.Response(c, "Please enter your email or password", 400, "bad request body", nil)
		return
	}


	user, err := u.Repository.FindUserByEmail(loginRequest.Email)
	if err != nil {
		util.Response(c, "user does not exist", 404, "user not found", nil)
		return
	}
	if user.LoginCounter >= 3 {
		user.IsLocked = true
		user.UpdatedAt = time.Now()
		err = u.Repository.UpdateUser(user)
		if err != nil {
			return
		}
		util.Response(c, "Your account has been lock after 3 failed attempt, contact customer care for assistance", 200, "success", nil)
		return
	}

	hashPass, err := util.HashPassword(user.Password)
	if err != nil {
		util.Response(c, "could not hash password", 500, "internal server error", nil)
		return
	}

	user.Password = hashPass

	
	if user.Password != loginRequest.Password {
		user.LoginCounter++
		err := u.Repository.UpdateUser(user)
		if err != nil {
			util.Response(c, "internal server error", 500, "user not found", nil)
			return
		}
		util.Response(c, "password mismatch", 404, "user not found", nil)
		return
	}

	//Generate token
	accessClaims, refreshClaims := middleware.GenerateClaims(user.Email)

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
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil)
}
