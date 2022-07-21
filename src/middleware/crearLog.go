package middleware

import (
	"log"
	"net/http"
	"time"
)

// Este Middleware genera un Log de la petición que se ejecuta
func CrearLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s -> %s %s: %s", r.RemoteAddr, r.Method, r.URL, time.Since(t))
	})
}
