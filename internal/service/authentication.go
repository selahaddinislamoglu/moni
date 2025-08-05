package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/selahaddinislamoglu/moni/internal/model"
)

type AuthenticationService struct {
	secretService Secret
}

func NewAuthenticationService() Authentication {
	return &AuthenticationService{}
}

func (s *AuthenticationService) SetupSecretService(secret Secret) {
	s.secretService = secret
}

func (s *AuthenticationService) Login(request model.LoginRequest) (*model.LoginResponse, error) {

	if request.Username != "admin" || request.Password != "password" {
		return nil, errors.New("invalid credentials")
	}

	claims := &model.JWTClaims{
		Username: request.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "moni",
			Subject:   request.Username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.secretService.getJWTsecret())
	if err != nil {
		fmt.Println("Failed to sign token:", err)
		return nil, err
	}
	return &model.LoginResponse{
		Token: tokenString,
	}, nil
}
