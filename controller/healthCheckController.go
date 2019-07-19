package controller

import (
	"fmt"
	"net/http"
)


func HealthCheck(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("health check is ok!")
	w.Write([]byte("health check is ok!"))
}
