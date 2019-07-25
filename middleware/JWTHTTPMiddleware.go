package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"sparta/config"
	"sparta/core"
	"strconv"
)

func JWTHTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			core.ResultFail(w, "Token is not valid")
			return
		}

		// 解析 JWT 签名，是否有效 赋值于 token.Valid
		SecretKey := config.Viper.GetString("secret_key")
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

		if err == nil {
			if token.Valid {
				// 取出 id 放在 request上
				claims := token.Claims.(jwt.MapClaims)
				id := strconv.FormatInt(int64(claims["id"].(float64)), 10)
				r.Header.Add("id", id)
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				core.ResultFail(w, "Token is not valid")
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			core.ResultFail(w, "Unauthorized access to this resource")
		}

	})
}

