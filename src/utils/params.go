package utils

import (
	"net/http"
	"strings"
)

// Obtiene el QueryParam que se le indique y lo retorna en String (si no existe retorna "")
func GetQueryParam(r *http.Request, param string) string {
	if !r.URL.Query().Has(param) {
		return ""
	}
	return r.URL.Query().Get(param)
}

// Obtiene el ultimo PathVariable de la Ruta (de /users/1, obtiene el 1)
// prefix es el /prefix/{valor} (incluyendo los "/" de delante y detrás)
func GetLastPathVariable(r *http.Request, prefix string) string {
	return strings.TrimPrefix(r.URL.Path, prefix)
}
