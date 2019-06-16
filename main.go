package main

import (
	"net/http"
	"sparta/config"
	_ "sparta/controller"
	"sparta/core"
	"time"
)

func main() {
	config.InitConfig()
	core.ConnectDB()

	server := &http.Server{
		Addr: ":5005",
		Handler: core.Router,
		ReadTimeout: 5 * time.Second,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
