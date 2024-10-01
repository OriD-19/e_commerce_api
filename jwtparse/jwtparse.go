package jwtparse

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

func ExtractTokensFromHeader(headers http.Header) string {
	authHeader := headers.Get("Authorization")

	if authHeader == "" {
		return ""
	}

	// index[1] contains the jwt
	splitToken := strings.Split(authHeader, "Bearer ")

	if len(splitToken) != 2 {
		// error parsing the header
		return ""
	}

	return splitToken[1]
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// TODO Find a way to store the secret
		secret := "ULTRA MEGA SECRET STRING"
		return []byte(secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is not valid - unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("claims of unauthorized token are not valid")
	}

	return claims, nil
}
