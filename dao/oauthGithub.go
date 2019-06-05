package dao

import (
	"fmt"
	"sparta/core"
	"sparta/model"
)

type OAuthGithubDao struct {}

func (this *OAuthGithubDao) Insert(githubUser *model.GithubUser) int64 {
	result, err := core.DB.Exec(
		"INSERT INTO OAuthGithub values (?,?,?,?,?,?,?,?,?)",
		githubUser.Id, githubUser.Name, githubUser.Avatar, githubUser.Company, githubUser.Blog, githubUser.Email, githubUser.Location, githubUser.Create_time, githubUser.Update_time)

	if err != nil {
		fmt.Println(err)
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return id
}


func (this *OAuthGithubDao) SelectUserById(id int64) model.GithubUser {
	fmt.Println(id)

	var user model.GithubUser
	err := core.DB.QueryRow("SELECT * FROM OAuthGithub").Scan(&user.Id, &user.Name, &user.Avatar)
	if err != nil {
		fmt.Println(err)
	}

	return user
}