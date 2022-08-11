package middleware

import (
	"myAPI/src/logger"
	"net/http"
	"time"
)

var mlog = logger.Logger

// Este Middleware genera un Log de la petición que se ejecuta
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next.ServeHTTP(w, r)
		mlog.Printf("%s -> %s %s : %s", r.RemoteAddr, r.Method, r.URL, time.Since(t))
	})
}
