package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/selahaddinislamoglu/moni/internal/service"
)

type AuthorizationController struct {
	authorizationService service.Authorization
}

func NewAuthorizationController() *AuthorizationController {
	return &AuthorizationController{}
}

func (a *AuthorizationController) Setup(authorizationService service.Authorization) {
	a.authorizationService = authorizationService
}

func (a *AuthorizationController) IsAuthorized(ctx *gin.Context) bool {

	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		ctx.JSON(401, gin.H{"error": "Authorization header is missing"})
		return false
	}

	token := tokenString[len("Bearer "):]
	if token == "" {
		ctx.JSON(401, gin.H{"error": "Token is missing"})
		return false
	}

	authorized := a.authorizationService.IsAuthorized(token)
	if !authorized {
		ctx.JSON(403, gin.H{"error": "Forbidden"})
	}
	return authorized
}
