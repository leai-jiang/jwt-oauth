package controller

import (
	"net/http"
	"sparta/core"
)

const (
	CLIENT_ID = "eed6beabf09a8713d3a7"
	//CLIENT_SECRET = "801d470a630d99d2a6d6a0c05af875f326b3f9d5"
	code = "3320e5b108ea531e9ff848f67eaf2b8429b570fd"
)

var OAuth = new(OAuthController)

type OAuthController struct {}

func (* OAuthController) login(w http.ResponseWriter, r *http.Request) {
	redirectUrl := "https://github.com/login/oauth/authorize?client_id=" + CLIENT_ID + "&redirect_uri="
	http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
}

func init() {
	core.Router.Register("/api/login/github", OAuth.login)
}
