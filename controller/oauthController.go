package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sparta/core"
	"sparta/dao"
	"sparta/model"
	"time"
)

const (
	//ClientId = "eed6beabf09a8713d3a7"
	//ClientSecret = "801d470a630d99d2a6d6a0c05af875f326b3f9d5"
	ClientId = "ecf4d78a2de563fbf68a"
	ClientSecret = "01f41a42bfdd5564f4b6d7191c3d70d268f445cf"
)

var githubUserDao = new(dao.OAuthGithubDao)

var OAuth = new(OAuthController)

type OAuthController struct {}

func (* OAuthController) Login(w http.ResponseWriter, r *http.Request) {
	redirectUrl := "https://github.com/login/oauth/authorize?client_id=" + ClientId
	http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
}

func (* OAuthController) GetAccessToken(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Panic(err)
	}
	code := r.Form.Get("code")
	requestUrl := "https://github.com/login/oauth/access_token?client_id=" + ClientId + "&client_secret=" + ClientSecret + "&code=" + code

	resp, err := http.Post(requestUrl, "application/x-www-form-urlencoded", nil)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	query, err := url.ParseQuery(string(body))

	accessToken := query["access_token"][0]
	log.Println(`token ` + accessToken)


	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", `token ` + accessToken)

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {

	}

	var githubUser *model.GithubUser
	if err := json.Unmarshal(data, &githubUser); err != nil {

	} else {
		if rows := len(githubUserDao.SelectUserById(githubUser.Id)); rows == 0 {
			githubUser.Create_time = time.Now()
			githubUser.Update_time = time.Now()
			githubUserDao.Insert(githubUser)
		}
	}


	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

func init() {
	core.Router.Register("/api/login/github", OAuth.Login)
	core.Router.Register("/api/oauth/redirect", OAuth.GetAccessToken)
}
