package handlers

import (
	"fmt"
	"myAPI/src/utils"
	"net/http"
)

// Responde un "Hola Mundo!" en formato JSON
func HolaMundo(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		utils.RJSON(w, http.StatusOK, utils.JSON{
			"message": "Hola Mundo!",
		})

	case "POST":
		json := utils.JSON{}
		if err := utils.LJSON(w, r, &json); err != nil {
			return
		}

		utils.RJSON(w, http.StatusOK, utils.JSON{
			"message": fmt.Sprint("El mensaje que enviaste:", json["message"]),
		})

	default:
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": "Las peticiones a esta ruta deben ser GET o POST, no se permite: " + r.Method,
		})
	}
}
