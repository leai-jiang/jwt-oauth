package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sparta/config"
	"sparta/core"
	"sparta/dao"
	"sparta/model"
	"strconv"
	"strings"
	"time"
)

var githubUserDao = new(dao.OAuthGithubDao)

var OAuth = new(OAuthController)

type OAuthController struct {}

func (* OAuthController) Login(w http.ResponseWriter, r *http.Request) {
	var clientId string

	if isProd := strings.Compare("production", os.Getenv("env")) == 0; isProd {
		clientId = config.Viper.GetString("oauth_github_prod.clientId")
	} else {
		clientId = config.Viper.GetString("oauth_github_dev.clientId")
	}
	redirectUrl := "https://github.com/login/oauth/authorize?client_id=" + clientId
	http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
}

func (* OAuthController) GetAccessToken(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Panic(err)
	}
	code := r.Form.Get("code")

	var clientId string
	var clientSecret string

	if isProd := strings.Compare("production", os.Getenv("env")) == 0; isProd {
		clientId = config.Viper.GetString("oauth_github_prod.clientId")
		clientSecret = config.Viper.GetString("oauth_github_prod.clientSecret")
	} else {
		clientId = config.Viper.GetString("oauth_github_dev.clientId")
		clientSecret = config.Viper.GetString("oauth_github_dev.clientSecret")
	}

	requestUrl := "https://github.com/login/oauth/access_token?client_id=" + clientId + "&client_secret=" + clientSecret + "&code=" + code

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

	cookie := &http.Cookie{
		Name: "u_t",
		Value: "github@" + strconv.FormatInt(githubUser.Id, 10),
		Path: "/",
		HttpOnly: true,
		MaxAge: int(time.Hour * 24 / time.Second),
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "http://localhost:3000", http.StatusTemporaryRedirect)
}

func init() {
	core.Router.Register("/api/login/github", OAuth.Login)
	core.Router.Register("/api/oauth/redirect", OAuth.GetAccessToken)
}
