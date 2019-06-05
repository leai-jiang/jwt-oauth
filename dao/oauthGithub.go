package dao

import (
	"fmt"
	"log"
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


func (this *OAuthGithubDao) SelectUserById(id int64) []model.GithubUser {
	rows,err := core.DB.Query("SELECT * FROM OAuthGithub WHERE id = ?",id)
	if err != nil {
		log.Println(err)
		return nil
	}
	var users []model.GithubUser
	for rows.Next() {
		var user model.GithubUser
		err := rows.Scan(&user.Id,&user.Name)
		if err != nil{
			log.Println(err)
			continue
		}
		users = append(users,user)
	}
	rows.Close()
	return users
}