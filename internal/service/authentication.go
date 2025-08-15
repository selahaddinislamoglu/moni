package service

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
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

func (s *AuthenticationService) Setup(secret Secret) {
	s.secretService = secret
}

func (s *AuthenticationService) verifyUserCredentials(username, password string) bool {

	cmd := exec.Command("su", "-c", "echo AUTH_OK", username)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Error creating stdin pipe:", err)
		return false
	}
	defer stdin.Close()

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Start()
	if err != nil {
		fmt.Println("Failed to start su command:", err)
		return false
	}

	_, err = stdin.Write([]byte(password + "\n"))
	if err != nil {
		fmt.Println("Failed to write password:", err)
		return false
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Authentication failed")
		fmt.Println("stderr:", stderr.String())
		return false
	}

	output := stdout.String()
	if output == "AUTH_OK\n" {
		fmt.Println("Authentication successful")
	} else {
		fmt.Println("Authentication failed")
		fmt.Println("Output:", output)
		fmt.Println("stderr:", stderr.String())
		return false
	}
	return true
}

func (s *AuthenticationService) Login(request model.LoginRequest) (*model.LoginResponse, error) {

	if !s.verifyUserCredentials(request.Username, request.Password) {
		return nil, errors.New("invalid username or password")
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
