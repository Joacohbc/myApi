package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"myAPI/src/handlers"
	"myAPI/src/middleware"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

const (
	// Puerto predeterminado si no se selecciona uno
	defaultPort     string = "8080"
	defaultFrontEnd string = "../static"
	version         string = "v1.0"
)

var (
	portSelected string
	frontEndDir  string = "../static"
	buff         bytes.Buffer
)

func init() {
	flag.StringVar(&portSelected, "port", defaultPort, fmt.Sprintf("Indica el puerto de escucha del servidor (predeterminadamente: %s)", defaultPort))
	flag.StringVar(&frontEndDir, "front", defaultFrontEnd, fmt.Sprintf("Indica el frontend que (predeterminadamente: %s)", defaultPort))
	flag.Parse()

	abs, err := filepath.Abs(frontEndDir)
	if err != nil {
		fmt.Println("Error al obtener la ruta del frontend:", err)
		os.Exit(1)
	}
	frontEndDir = abs
}

func main() {
	//Creo el Router de Rutas
	mux := http.NewServeMux()

	// Sirvo el Front-end
	mux.Handle("/", http.FileServer(http.Dir(frontEndDir)))

	// Simple endpoint para realizar peticiones POST y GET
	mux.Handle("/api", middleware.Logger(http.HandlerFunc(handlers.HolaMundo)))

	// Endpoint completo
	mux.Handle("/users/", middleware.Logger(http.HandlerFunc(handlers.Personas)))

	//Personalizo el servidor
	s := &http.Server{
		Addr:           ":" + portSelected,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 10 << 20,
	}

	wait := make(chan struct{})
	// Inicio el servidor
	go func() {
		// Creo un canal os.Signal
		exit := make(chan os.Signal, 1)

		// Que cuando reciba una señal de cierre que siga con el código
		signal.Notify(exit, os.Interrupt)
		<-exit

		// Cierro el servidor
		fmt.Println("Cerrando servidor...")

		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*15))
		defer cancel()
		if err := s.Shutdown(ctx); err != nil {
			fmt.Println("Error al cerrar el servidor: " + err.Error())
			os.Exit(1)
		}
		fmt.Println("Servidor cerrado con exito <3")

		close(wait)
	}()

	fmt.Print(`
	__  __          _    ____ ___   ____           _   
	|  \/  |_   _   / \  |  _ \_ _| |  _ \ ___  ___| |_ 
	| |\/| | | | | / _ \ | |_) | |  | |_) / _ \/ __| __|
	| |  | | |_| |/ ___ \|  __/| |  |  _ <  __/\__ \ |_ 
	|_|  |_|\__, /_/   \_\_|  |___| |_| \_\___||___/\__|
			|___/ `)
	fmt.Println("Version:", version)

	fmt.Println("- Puerto:", portSelected)
	fmt.Println("- Frontend:", frontEndDir)
	fmt.Println("- Incio:", time.Now().Format("01/02/2006 15:04:05"))

	fmt.Println("\nServidor escuchando...")
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Println("Error al iniciar el servidor:", err)
		os.Exit(1)
	}

	<-wait
}
