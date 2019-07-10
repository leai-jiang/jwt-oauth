package middleware

import (
	"encoding/json"
	"net/http"
	"sparta/jwt"
	"sparta/view"
)

var SecretKey = "It is a secret key"

func JWTHTTPMiddleware (h func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("token")

		token, err := new(jwt.Parse).Parse(tokenString, SecretKey)
		if err != nil {

		}
		if token != nil && token.Valid {
			h(w, r)
		}

		res := &view.BaseView{
			RetCode: -1,
			Data: nil,
			Message: err.Error(),
		}

		ress, _ := json.Marshal(res)

		w.Write(ress)
	}
}




