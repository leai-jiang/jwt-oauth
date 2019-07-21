package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)


func LoggerMiddleware(next http.Handler) http.Handler {
	logger := log.New()

	logger.Out = os.Stdout

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.WithFields(log.Fields{
			"path": r.URL.Path,
			"method": r.Method,
			"host": r.Host,
		}).Info("[sparta-log]")

		next.ServeHTTP(w, r)
	})
}


