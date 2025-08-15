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
		tokenString = ctx.Query("token")
		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "Authorization or token query parameter is missing"})
			return false
		}
	} else {
		tokenString = tokenString[len("Bearer "):]
	}

	authorized := a.authorizationService.IsAuthorized(tokenString)
	if !authorized {
		ctx.JSON(403, gin.H{"error": "Forbidden"})
	}
	return authorized
}
