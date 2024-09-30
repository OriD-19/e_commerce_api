package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func JWTAuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := extractTokensFromHeader(r.Header)

		if tokenString == "" {
			http.Error(w, "no token provided", http.StatusUnauthorized)
			next.ServeHTTP(w, r)
		}

		claims, err := parseToken(tokenString)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			next.ServeHTTP(w, r)
		}

		fmt.Println(claims["expires"])
		expires := int64(claims["expires"].(float64))

		if time.Now().Unix() > expires {
			http.Error(w, "token expired", http.StatusUnauthorized)
			next.ServeHTTP(w, r)
		}

		// After all the validations, pass to the next handler
		// pass it to the team ;)
		next.ServeHTTP(w, r)
	})
}

func extractTokensFromHeader(headers http.Header) string {
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

func parseToken(tokenString string) (jwt.MapClaims, error) {
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
