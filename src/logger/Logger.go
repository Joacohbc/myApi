package logger

import (
	"fmt"
	"io"
	"log"
	"myAPI/src/env"
	"os"
	"path/filepath"
	"time"
)

var (
	Logger  *log.Logger
	AbsPath string
)

func init() {

	if !env.GeneraLogFile {
		Logger = log.New(os.Stdout, "INFO: ", log.Lshortfile|log.Ldate|log.Ltime|log.Lmicroseconds)
		return
	}

	// Creo el directorio de logs
	if err := os.MkdirAll("./logs", 0744); err != nil {
		fmt.Println("Error al crear el directorio de logs: " + err.Error())
		os.Exit(1)
	}

	// 01 de Enero (02) 15:04:05 (03:04:05) 2006 (06)
	filename := "./logs/log_" + time.Now().Format("01-02-2006_15-04-05") + ".txt"

	// Creo el archivo de guardado
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0744)
	if err != nil {
		fmt.Println("Error al crear el archivo de logs: " + err.Error())
		os.Exit(1)
	}

	// Busco la ruta absoluta del archivo logs
	abs, err := filepath.Abs(filename)
	if err != nil {
		fmt.Println("Error al hallar la ruta absoluta del archivo de logs: " + err.Error())
		os.Exit(1)
	}

	AbsPath = abs

	// Creo un multiwriter con al salida consola (os.Stdout) y al archivo de logs
	mw := io.MultiWriter(os.Stdout, file)

	// Y le asigno el Writter al logger
	Logger = log.New(mw, "INFO: ", log.Lshortfile|log.Ldate|log.Ltime|log.Lmicroseconds)
}
