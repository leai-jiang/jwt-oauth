package controller

import (
	"net/http"
	"sparta/core"
)

func SignOut(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err == nil {
		token.MaxAge = 0
		http.SetCookie(w, token)
	}

}

func init() {
	core.Router.Register("/api/signout", SignOut)
}
