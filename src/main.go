package main

import (
	"log"
	"myAPI/src/handlers"
	"myAPI/src/middleware"
	"net/http"
	"time"
)

const (
	portSelected = "8080"
	staticFiles  = "./static"
)

func main() {

	//Creo el Router de Rutas
	router := http.NewServeMux()

	// Sirvo el Front-end
	router.Handle("/", http.FileServer(http.Dir(staticFiles)))

	// Simple endpoint para realizar peticiones POST y GET
	router.Handle("/api", middleware.CrearLog(http.HandlerFunc(handlers.HolaMundo)))

	// Endpoint completo
	router.Handle("/users/", middleware.CrearLog(http.HandlerFunc(handlers.Personas)))

	//Personalizo el servidor
	s := &http.Server{
		Addr:           ":" + portSelected,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 10 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
