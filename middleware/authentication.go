package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"orid19.com/ecommerce/api/jwtparse"
	"orid19.com/ecommerce/api/types"
)

func JWTAuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := jwtparse.ExtractTokensFromHeader(r.Header)

		if tokenString == "" {
			http.Error(w, "no token provided", http.StatusUnauthorized)
			next.ServeHTTP(w, r)
		}

		claims, err := jwtparse.ParseToken(tokenString)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			next.ServeHTTP(w, r)
		}

		fmt.Printf("Original type of claims: %T", claims["expires"])
		expires := int64(claims["expires"].(float64))

		if time.Now().Unix() > expires {
			http.Error(w, "token expired", http.StatusUnauthorized)
			next.ServeHTTP(w, r)
		}

		var contextUser types.ContextUser = "user_id"

		ctxWithUser := context.WithValue(r.Context(), contextUser, claims["user_id"])
		requestWithUser := r.WithContext(ctxWithUser)

		fmt.Println("This is the value of the user:", claims["user_id"])

		// After all the validations, pass to the next handler
		// pass it to the team ;)
		// it is contained in the context of the request
		next.ServeHTTP(w, requestWithUser)
	})
}
