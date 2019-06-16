package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"sparta/core"
	"sparta/dao"
	"strconv"
	"strings"
)

var githubUserDao1 = new(dao.OAuthGithubDao)

type loginController struct {}

var login = new(loginController)

type responseBody struct {
	RetCode int64
	Data interface{}
	message string
}

func (this *loginController) checkLogin(w http.ResponseWriter, r *http.Request)  {
	body := new(responseBody)

	token, err := r.Cookie("u_t")
	if err != nil {
		if err == http.ErrNoCookie {
			body = &responseBody{
				RetCode: 0,
				Data: nil,
				message: "Did not logged",
			}
		}
	}

	if token == nil {
		body = &responseBody{
			RetCode: 0,
			Data: nil,
			message: "Did not logged",
		}
	} else {
		if has := strings.HasPrefix(token.Value, "github"); has {
			str := strings.Split(token.Value, "@")[1]
			id, _ := strconv.ParseInt(str, 10, 64)
			userInfo := githubUserDao1.SelectUserById(id)
			if userInfo != nil {
				body = &responseBody{
					RetCode: 0,
					Data: userInfo,
					message: "logged",
				}
			}
		}
	}

	rb, err := json.Marshal(body)
	if err != nil {
		log.Panic("Cannot transfer body to json")
	}
	w.Write(rb)
}

func init() {
	core.Router.Register("/api/login", login.checkLogin)
}
