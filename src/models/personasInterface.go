package models

import (
	"encoding/json"
	"fmt"
	"myAPI/src/logger"
	"net/http"
	"os"
	"sync"
)

// Variable de mi archivo para realizar Logs
var mlog = logger.Logger

// Ruta donde guardar el archivos de People
const personasJsonPath string = "./people.json"

// Para evitar que se escriba y lea al mismo tiempo el archivo
var mux sync.Mutex

type PeopleInterface interface {
	ObtenerPersonas(w http.ResponseWriter) (map[int]People, error)
	ActualizarPersonas(persona map[int]People) error
}

func PeopleService() PeopleInterface {
	return peopleFuncs{}
}

type peopleFuncs struct{}

// Retorna un Map de personas y una Flag que indica si se realizo con éxito la lectura de personas (true)
// o si ocurrió un error (false)
func (p peopleFuncs) ObtenerPersonas(w http.ResponseWriter) (map[int]People, error) {

	mux.Lock()
	defer mux.Unlock()

	// Leo el archivo
	b, err := os.ReadFile(personasJsonPath)
	if err != nil {
		mlog.Println("Error al leer el archivo de personas:", err)
		return map[int]People{}, fmt.Errorf("error al intentar leer las personas, inténtelo mas tarde")
	}

	// Lo cargo en un map de persona
	var personas map[int]People
	if err = json.Unmarshal(b, &personas); err != nil {
		mlog.Println("Error al cargar el archivo de personas al map de personas:", err)
		return map[int]People{}, fmt.Errorf("error al intentar leer las personas, inténtelo mas tarde")
	}

	// Y lo retorno
	return personas, nil
}

// Sobrescribe el archivo json de las personas
func (p peopleFuncs) ActualizarPersonas(persona map[int]People) error {

	mux.Lock()
	defer mux.Unlock()

	b, err := json.MarshalIndent(persona, " ", "\t")
	if err != nil {
		mlog.Println("Error al actualizar el archivo de personas:", err)
		return err
	}

	err = os.WriteFile(personasJsonPath, b, 0644)
	if err != nil {
		mlog.Println("Error al actualizar el archivo de personas:", err)
		return err
	}

	return nil
}

func init() {
	// Si el archivo existe de las personas, retorno para continuar con el código
	if _, err := os.Stat(personasJsonPath); err == nil {
		return
	}

	// Sino existe el archivo, creo un nuevo JSON vació
	b, err := json.MarshalIndent(map[int]People{}, " ", "\t")
	if err != nil {
		mlog.Println("Error al crear el archivo de persona:", err)
	}

	// Y creo el archivo
	err = os.WriteFile(personasJsonPath, b, 0644)
	if err != nil {
		mlog.Println("Error al crear el archivo de persona:", err)
	}
}
