package controller

import (
	"net/http"
	"sparta/core"
	"strconv"
)

type UserController struct {}

func (u *UserController) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	idStr := r.Header.Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		core.ResultFail(w, err.Error())
		return
	}
	user, err := githubUserDao.SelectUserById(id)
	if err != nil {
		core.ResultFail(w, err.Error())
		return
	}

	core.ResultOk(w, user)
}
