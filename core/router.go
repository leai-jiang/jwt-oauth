package core

import (
	"fmt"
	"net/http"
)

var Router = new(RouteHandler)

type RouteHandler struct {}

var mux = make(map[string]func(http.ResponseWriter, *http.Request))

func (this *RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path
	fmt.Println(requestPath)

	if handler, ok := mux[requestPath]; ok {
		handler(w, r)
		return
	}

	http.Error(w, "URL Not Found:" + requestPath, http.StatusBadRequest)
}

func (this *RouteHandler) Register(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	mux[pattern] = handler
}

