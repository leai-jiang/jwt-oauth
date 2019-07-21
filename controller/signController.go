package controller

import (
	"net/http"
	"time"
)

type SignController struct {}

func (s *SignController) SignOut(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("token")

	if err == nil {
		c := http.Cookie{
			Name:     "token",
			Path:     "/",
			Value:    "",
			HttpOnly: true,
			Expires:  time.Now().AddDate(0,0,-1),
			MaxAge:   0,
		}
		http.SetCookie(w, &c)
	}
}

func (s *SignController) SignIn(w http.ResponseWriter, r *http.Request) {

}

func (s *SignController) Session(w http.ResponseWriter, r *http.Request) {

}
