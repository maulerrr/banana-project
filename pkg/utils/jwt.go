package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/maulerrr/banana/pkg/models"
	"time"
)

func GenerateToken(user models.User, jwtKey string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)

	claims := &models.Claims{
		ID:       user.UserID,
		Email:    user.Email,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))

	return tokenString, err
}

func ParseToken(tokenString string, jwtKey string) (*models.Claims, error) {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
