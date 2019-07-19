package middleware

import (
	"net/http"
)

var SecretKey = "It is a secret key"

func JWTHTTPMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//tokenString := r.Header.Get("token")
		//
		//token, err := new(jwt.Parse).Parse(tokenString, SecretKey)
		//if err != nil {
		//
		//}

	})
}

