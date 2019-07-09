package middleware

import (
	"encoding/json"
	"net/http"
	"sparta/jwt"
)

type Response struct {
	RetCode 	int32 		`json:"retCode"`
	Data 		interface{} `json:"data"`
	Message 	string		`json:"message"`
}

func WrapJWTHTTPMiddleware(key string) func(func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func (h func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("token")

			token, err := new(jwt.Parse).Parse(tokenString, key)
			if err != nil {

			}
			if token != nil && token.Valid {
				h(w, r)
			}

			res := &Response{
				RetCode: -1,
				Data: nil,
				Message: "permission not allow",
			}

			ress, e := json.Marshal(res)
			if e != nil {}
			w.Write(ress)
		}
	}
}




