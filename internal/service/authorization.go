package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/selahaddinislamoglu/moni/internal/model"
)

type AuthorizationService struct {
	secretService Secret
}

func NewAuthorizationService() Authorization {
	return &AuthorizationService{}
}

func (a *AuthorizationService) Setup(secret Secret) {
	a.secretService = secret
}

func (a *AuthorizationService) IsAuthorized(token string) bool {
	parser := new(jwt.Parser)

	claims := &model.JWTClaims{}

	parsed, err := parser.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return a.secretService.getJWTsecret(), nil
	})

	if err != nil || !parsed.Valid {
		return false
	}

	return true
}
