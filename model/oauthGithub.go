package model

import "time"

type GithubUser struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Company     string `json:"company"`
	Blog        string `json:"blog"`
	Email       string `json:"email"`
	Location    string `json:"location"`
	Create_time time.Time `json:"createTime"`
	Update_time time.Time `json:"updateTime"`
}