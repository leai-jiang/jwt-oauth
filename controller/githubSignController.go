package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sparta/config"
	"sparta/dao"
	"sparta/model"
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

//
func (g *GithubSignController) SignIn(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Panic(err)
	}
	code := r.Form.Get("code")

	accessToken, err := g.getAccessToken(code)
	if err != nil {

	}

	user, err := g.getUserInfo(accessToken)

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

	accessToken = query["access_token"][0]
	log.Println(`token ` + accessToken)

	return accessToken, nil
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

