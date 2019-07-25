package controller

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sparta/config"
	"sparta/core"
	"sparta/dao"
	"sparta/model"
	"time"
)

const (
	GithubLoginUrl = "https://github.com/login/oauth/authorize"

	GithubAccessTokenApi = "https://github.com/login/oauth/access_token"

	GithubUserApiUrl = "https://api.github.com/user"
)

var githubUserDao = new(dao.OAuthGithubDao)

type GithubSignController struct {}

// 跳转到 github 登录界面
func (g *GithubSignController) RedirectToGithub(w http.ResponseWriter, r *http.Request) {
	clientId := config.Viper.GetString("oauth_github_dev.clientId")
	redirectUrl := GithubLoginUrl + "?client_id=" + clientId
	http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
}

// github OAuth 登陆
func (g *GithubSignController) SignIn(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Panic(err)
	}
	code := r.Form.Get("code")

	accessToken, err := g.getAccessToken(code)
	if err != nil {
		core.ResultFail(w, err.Error())
		return
	}

	user, err := g.getUserInfo(accessToken)
	if err != nil {
		core.ResultFail(w, err.Error())
		return
	}

	// 查数据库，有即更新，无则保存
	alreadyExistUser, err := githubUserDao.SelectUserById(user.Id)
	if err != nil {
		fmt.Println("The user is new", user)
		githubUserDao.Insert(user)
	}
	fmt.Println("The user already exist", alreadyExistUser)

	// JWT 签名
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1))
	claims["iat"] = time.Now().Unix()
	claims["id"] = user.Id

	token.Claims = claims

	SecretKey := config.Viper.GetString("secret_key")
	tokenString, err := token.SignedString([]byte(SecretKey))

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: tokenString,
		Path: "/",
		Expires: time.Now().Add(time.Hour * time.Duration(1)),
		HttpOnly: true,
	})

	homePage := config.Viper.GetString("home_page_dev")
	http.Redirect(w, r, homePage, http.StatusTemporaryRedirect)
}

// 获取 access_token
func (g *GithubSignController) getAccessToken(code string) (accessToken string, err error) {
	clientId := config.Viper.GetString("oauth_github_dev.clientId")
	clientSecret := config.Viper.GetString("oauth_github_dev.clientSecret")

	requestUrl := GithubAccessTokenApi + "?client_id=" + clientId + "&client_secret=" + clientSecret + "&code=" + code

	resp, err := http.Post(requestUrl, "application/x-www-form-urlencoded", nil)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	query, err := url.ParseQuery(string(body))
	if err != nil {
		return "", err
	}

	if len(query["access_token"]) > 0 {
		accessToken = query["access_token"][0]
	}
	log.Println(`token ` + accessToken)

	return
}

// 拉取 github 用户信息
func (* GithubSignController) getUserInfo(accessToken string) (*model.GithubUser, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", GithubUserApiUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", `token ` + accessToken)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var githubUser *model.GithubUser
	if err := json.Unmarshal(data, &githubUser); err != nil {
		return nil, err
	}

	return githubUser, nil
}

