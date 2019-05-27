package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"sparta/core"
	_ "sparta/controller"
)

func main() {
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
