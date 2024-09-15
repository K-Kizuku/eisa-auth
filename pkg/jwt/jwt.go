package jwt

import (
	"errors"
	"fmt"

	env "github.com/K-Kizuku/eisa-auth/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(env.JWTSecret)
	if err != nil {
		return "", err
	}

	fmt.Println(tokenString)
	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return env.JWTSecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["user_id"].(string), nil
	} else {
		return "", err
	}
}
