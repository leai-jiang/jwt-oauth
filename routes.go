package main

import (
	"net/http"
	"sparta/controller"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
}

type Routes []Route

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

var signController = new(controller.SignController)
var githubSignController = new(controller.GithubSignController)

var routes = Routes{
	Route{"HealthCheck", "GET", "/health", controller.HealthCheck},

	// 登录
	Route{"SignIn", "POST", "/login", signController.SignIn},

	// github 第三方登录
	Route{ "GithubSignIn", "POST", "/login/github", githubSignController.RedirectToGithub},

	// 登出
	Route{"SignOut", "POST", "/logout", signController.SignOut},
}

var restRoutes = Routes{
	Route{},
}
