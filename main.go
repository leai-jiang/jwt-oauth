package main

import (
	"fmt"
	"log"
	"net/http"
	"sparta/controller"
	_ "sparta/controller"
	"sparta/core"
	"time"
)

func main() {
	core.ConnectDB()

	core.Router.Register("/api/login/github", controller.OAuth.Login)
	core.Router.Register("/api/oauth/redirect", controller.OAuth.GetAccessToken)

	server := &http.Server{
		Addr: ":5005",
		Handler: core.Router,
		ReadTimeout: 5 * time.Second,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
		log.Panic(err)
	}
}
