package model

import "time"

type GithubUser struct {
	Id          int64 `json:"id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar_url"`
	Company     string `json:"company"`
	Blog        string `json:"blog"`
	Email       string `json:"email"`
	Location    string `json:"location"`
	Create_time time.Time `json:"createTime"`
	Update_time time.Time `json:"updateTime"`
}