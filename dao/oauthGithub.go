package dao

import (
	"log"
	"sparta/core"
	"sparta/model"
	"time"
)

type OAuthGithubDao struct {}

func (o *OAuthGithubDao) Insert(githubUser *model.GithubUser) int64 {
	result, err := core.DB.Exec(
		"INSERT INTO OAuthGithub values (?,?,?,?,?,?,?,?,?)",
		githubUser.Id, githubUser.Name, githubUser.Avatar, githubUser.Company, githubUser.Blog, githubUser.Email, githubUser.Location, time.Now(), time.Now())

	if err != nil {
		log.Println(err)
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0
	}
	return id
}

func (o *OAuthGithubDao) Update(githubUser *model.GithubUser) int64 {
	result, err := core.DB.Exec(
		"UPDATE OAuthGithub set (`name`, `avatar_url`, `company`, `blog`, `email`, `location`) = (?,?,?,?,?,?,?,) WHERE id = ?",
		githubUser.Name, githubUser.Avatar, githubUser.Company, githubUser.Blog, githubUser.Email, githubUser.Location,
		githubUser.Id,
	)
	if err != nil {
		log.Println(err)
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0
	}
	return id
}


func (o *OAuthGithubDao) SelectUserById(id int64) (*model.GithubUser, error) {
	var user = new(model.GithubUser)
	err := core.DB.QueryRow("SELECT `id`, `name`, `avatar` FROM OAuthGithub WHERE id = ?",id).Scan(&user.Id, &user.Name, &user.Avatar)

	if err != nil {
		return nil, err
	}

	return user, nil
}
