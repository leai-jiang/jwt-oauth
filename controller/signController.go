package controller

import (
	"net/http"
	"time"
)

type SignController struct {}

func (s *SignController) SignOut(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err == nil {
		token.Expires = time.Now().AddDate(0,0,-1)
		http.SetCookie(w, token)
	}
}

func (s *SignController) SignIn(w http.ResponseWriter, r *http.Request) {

}

func (s *SignController) Session(w http.ResponseWriter, r *http.Request) {

}
