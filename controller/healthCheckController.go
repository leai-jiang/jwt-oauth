package controller

import (
	"fmt"
	"log"
	"net/http"
	"sparta/core"
)

var healthCheckController = new(HealthCheckController)

type HealthCheckController struct {}

func (p *HealthCheckController) check(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("health check is ok!")
	_, err := w.Write([]byte("health check is ok!"))

	if err != nil {
		log.Panic(err)
	}
}

func init() {
	core.Router.Register("/health-check", healthCheckController.check)
}