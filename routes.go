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
var userController = new(controller.UserController)

var routes = Routes{
	Route{
		"HealthCheck",
		"GET",
		"/health",
		controller.HealthCheck,
	},

	// 登录
	Route{
		"SignIn",
		"POST",
		"/login",
		signController.SignIn,
	},

	// github 第三方登录
	Route{
		"GithubSignIn",
		"GET",
		"/login/github",
		githubSignController.RedirectToGithub,
	},

	// github 第三方登陆成功回调
	Route{
		"GithubCallback",
		"GET",
		"/github/callback",
		githubSignController.SignIn,
	},

	// 登出
	Route{
		"SignOut",
		"POST",
		"/logout",
		signController.SignOut,
	},
}

var restRoutes = Routes{
	Route{
		"GetUserInfo",
		"POST",
		"/user",
		userController.GetUserInfo,
	},
}
