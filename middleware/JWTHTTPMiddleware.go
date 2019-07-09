package middleware

import (
	"net/http"
	"sparta/jwt"
)

func WrapJWTHTTPMiddleware(key string) func(func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func (h func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("token")

			if token, err := new(jwt.Parse).Parse(tokenString, key); err != nil {

			}


		}
	}
}




