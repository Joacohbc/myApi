package env

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// Indica la version de myApi
const Version string = "v1.1"

var (
	PortSelected  string
	FrontEndDir   string
	GeneraLogFile bool
)

const (
	// Puerto predeterminado si no se selecciona uno
	defaultPort          string = "8080"
	defaultFrontEnd      string = "./static"
	defaultGenereLogFile bool   = true
)

func init() {

	flag.StringVar(&PortSelected, "port", defaultPort, fmt.Sprintf("Indica el puerto de escucha del servidor (predeterminadamente: %s)", defaultPort))
	flag.StringVar(&FrontEndDir, "front", defaultFrontEnd, fmt.Sprintf("Indica el frontend que (predeterminadamente: %s)", defaultPort))
	flag.BoolVar(&GeneraLogFile, "log", defaultGenereLogFile, fmt.Sprintf("Determina si se generaran archivos de log o no (predeterminadamente: %v)", defaultGenereLogFile))
	flag.Parse()

	abs, err := filepath.Abs(FrontEndDir)
	if err != nil {
		fmt.Println("Error al obtener la ruta del frontend:", err)
		os.Exit(1)
	}
	FrontEndDir = abs
}
