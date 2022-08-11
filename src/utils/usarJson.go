package utils

import (
	"encoding/json"
	"io"
	"myAPI/src/logger"
	"net/http"
	"strconv"
)

var mlog = logger.Logger

// Abreviación de map[string]interface{}
type JSON map[string]interface{}

// Función que responde en JSON
func RJSON(w http.ResponseWriter, statusCode int, message interface{}) {

	jsonb, err := json.Marshal(message)
	if err != err {
		mlog.Println("Error al hacer marshal:", err)
		return
	}

	// Marco los headers
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(jsonb)))
	w.WriteHeader(statusCode)

	// Envió la respuesta
	w.Write(jsonb)
}

// Leer el JSON recibido de una Request, si la petición falla se envía un error
func LJSON[Any any](w http.ResponseWriter, r *http.Request, output Any) error {

	// Leo el Body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		mlog.Println("Error al leer el body de la peticion: ", err.Error())
		RJSON(w, http.StatusInternalServerError, JSON{
			"error": "Error al leer la peticion: " + err.Error(),
		})
		return err
	}

	// Y lo convierto en un JSON
	err = json.Unmarshal(body, output)
	if err != nil {
		mlog.Println("Error al convertir el Body a JSON: ", err.Error())
		RJSON(w, http.StatusBadRequest, JSON{
			"error": "Error al leer la peticion, formato erroneo: " + err.Error(),
		})
		return err
	}

	return nil
}
