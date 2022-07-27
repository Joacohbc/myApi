package main

import (
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
	router := http.NewServeMux()

	// Sirvo el Front-end
	router.Handle("/", http.FileServer(http.Dir(frontEndDir)))

	// Simple endpoint para realizar peticiones POST y GET
	router.Handle("/api", middleware.CrearLog(http.HandlerFunc(handlers.HolaMundo)))

	// Endpoint completo
	router.Handle("/users/", middleware.CrearLog(http.HandlerFunc(handlers.Personas)))

	//Personalizo el servidor
	s := &http.Server{
		Addr:           ":" + portSelected,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 10 << 20,
	}

	// Inicio el servidor
	go func() {
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
		if err := s.ListenAndServe(); err != nil {
			fmt.Println("Error al iniciar el servidor:", err)
			os.Exit(1)
		}
	}()

	// Creo un canal os.Signal
	exit := make(chan os.Signal, 1)

	// Que cuando reciba una señal de cierre que siga con el código
	signal.Notify(exit, os.Interrupt)
	<-exit

	// Cierro el servidor
	fmt.Println("Cerrando servidor...")
	if err := s.Close(); err != nil {
		fmt.Println("Error al cerrar el servidor: " + err.Error())
		os.Exit(1)
	}

	fmt.Println("Servidor cerrado con exito <3")
}
