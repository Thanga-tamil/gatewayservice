package service

import (
	"time"
	"gateway/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userId string) (string, error) {
	claims := jwt.MapClaims{ 
		"userId": userId, 
		"exp": time.Now().Add(time.Hour * 1).Unix(), 
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(utils.JWK)
	if err != nil {
		return "", err
	}

	return token, nil
}

