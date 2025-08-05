package service

import "crypto/rand"

type SecretService struct {
	secret []byte
}

func generateJWTSecret(length int) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func NewSecretService() Secret {
	secret, err := generateJWTSecret(32)
	if err != nil {
		panic("failed to generate JWT secret: " + err.Error())
	}
	return &SecretService{
		secret: secret,
	}
}

func (s *SecretService) getJWTsecret() []byte {
	return s.secret
}
