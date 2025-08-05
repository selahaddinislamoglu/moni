package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/selahaddinislamoglu/moni/internal/model"
	"github.com/selahaddinislamoglu/moni/internal/service"
)

type AuthenticationController struct {
	authenticationService service.Authentication
}

func NewAuthenticationController() *AuthenticationController {
	return &AuthenticationController{}
}

func (a *AuthenticationController) SetupAuthenticationService(authenticationService service.Authentication) {
	a.authenticationService = authenticationService
}

func (a *AuthenticationController) Login(ctx *gin.Context) {
	var request model.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	if request.Username == "" || request.Password == "" {
		ctx.JSON(400, gin.H{"error": "Username and password are required"})
		return
	}
	data, err := a.authenticationService.Login(request)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to login"})
		return
	}
	ctx.JSON(200, data)
}
