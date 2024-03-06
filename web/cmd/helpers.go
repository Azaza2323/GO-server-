package main

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
)

func ExtractUserIDFromToken(tokenStr string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("invalid token")
	}
	return claims.UserID, nil
}
