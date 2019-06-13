package middleware

import (
	"log"
	"net/http"
)

func Logger(hFunc func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(r.URL.Path, r.Header)
		hFunc(w, r)
	})
}
