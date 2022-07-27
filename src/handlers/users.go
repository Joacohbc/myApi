package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"myAPI/src/utils"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const personasJsonPath string = "./personas.json"

func init() {

	if _, err := os.Stat(personasJsonPath); err == nil {
		return
	}

	b, err := json.MarshalIndent(map[int]utils.People{}, " ", "\t")
	if err != nil {
		log.Println("Error al crear el archivo de persona:", err)
	}

	err = ioutil.WriteFile(personasJsonPath, b, 0644)
	if err != nil {
		log.Println("Error al crear el archivo de persona:", err)
	}
}

// Retorna un Map de personas
func obtenerPersonas() (map[int]utils.People, error) {

	b, err := ioutil.ReadFile(personasJsonPath)
	if err != nil {
		log.Println("Error al leer el archivo de personas:", err)
		return nil, err
	}

	var personas map[int]utils.People
	if err = json.Unmarshal(b, &personas); err != nil {
		log.Println("Error al leer el archivo de personas:", err)
		return nil, err
	}

	return personas, nil
}

// Sobrescribe el archivo json de las personas
func actualizarPersonas(persona map[int]utils.People) error {
	b, err := json.MarshalIndent(persona, " ", "\t")
	if err != nil {
		log.Println("Error al actualizar el archivo de personas:", err)
		return err
	}

	err = ioutil.WriteFile(personasJsonPath, b, 0644)
	if err != nil {
		log.Println("Error al actualizar el archivo de personas:", err)
		return err
	}

	return nil
}

// Endpoint - /users/ - Gestiona todas las peticiones
func Personas(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET", "HEAD":
		if strings.TrimPrefix(r.URL.Path, "/users/") == "" {
			getAllPeople(w, r)
		} else {
			getPerson(w, r)
		}

	case "POST":
		newPerson(w, r)

	case "DELETE":
		deletePerson(w, r)

	case "PUT", "PATCH":
		updatePerson(w, r)

	default:
		log.Println("Se hizo una peticion:", r.Method)
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": "Las peticiones a esta ruta deben ser GET, HEAD, POST, DELETE, PUT o PATCH",
		})
	}
}

// Endpoint - /users/ - GET/HEAD
func getPerson(w http.ResponseWriter, r *http.Request) {

	ci, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/users/"))
	if err != nil {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": "la persona pedida es invalida",
		})
		return
	}

	personaPedida := utils.People{
		CI: ci,
	}

	if err := personaPedida.ValidCI(); err != nil {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": err.Error(),
		})
		return
	}

	personas, err := obtenerPersonas()
	if err != nil {
		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": "Error al intentar leer los usuarios, intentelo mas tarde",
		})
		return
	}

	_, ok := personas[personaPedida.CI]
	if !ok {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": fmt.Sprintf("La persona con la CI %d no existe", personaPedida.CI),
		})
		return
	}

	utils.RJSON(w, http.StatusOK, personas[personaPedida.CI])
}

// Endpoint - /users/ - GET/HEAD
func getAllPeople(w http.ResponseWriter, _ *http.Request) {
	personasMap, err := obtenerPersonas()
	if err != nil {
		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": "Error al intentar leer los usuarios, intentelo mas tarde",
		})
		return
	}

	// Guardo los datos del map en una slice para enviarlos
	personas := []utils.People{}
	for _, v := range personasMap {
		personas = append(personas, v)
	}

	utils.RJSON(w, http.StatusOK, personas)
}

// Endpoint - /users/ - POST
func newPerson(w http.ResponseWriter, r *http.Request) {

	persona := utils.People{}
	if err := utils.LJSON(w, r, &persona); err != nil {
		return
	}

	if err := persona.ValidAll(); err != nil {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": err.Error(),
		})
		return
	}

	personas, err := obtenerPersonas()
	if err != nil {
		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": "Error al intentar leer los usuarios, intentelo mas tarde",
		})
		return
	}

	_, ok := personas[persona.CI]
	if ok {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": fmt.Sprintf("La persona con la CI %d ya existe", persona.CI),
		})
		return
	}

	personas[persona.CI] = persona
	if err := actualizarPersonas(personas); err != nil {
		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": "Error al agregar a la nueva persona a la lista, intentelo mas tarde",
		})
		return
	}

	utils.RJSON(w, http.StatusCreated, utils.JSON{
		"message": fmt.Sprintf("La persona %s %s fue creada con exito", persona.Name, persona.Surname),
	})
}

// Endpoint - /users/ - PUT/PATCH
func updatePerson(w http.ResponseWriter, r *http.Request) {

	newPerson := utils.People{}
	if err := utils.LJSON(w, r, &newPerson); err != nil {
		return
	}

	if err := newPerson.ValidAll(); err != nil {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": err.Error(),
		})
		return
	}

	personas, err := obtenerPersonas()
	if err != nil {
		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": "Error al intentar leer los usuarios, intentelo mas tarde",
		})
		return
	}

	_, ok := personas[newPerson.CI]
	if !ok {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": fmt.Sprintf("La persona con la CI %d no existe", newPerson.CI),
		})
		return
	}

	// personas[newPerson.CI].Birthdate -> Fecha de nacimiento actual de la persona
	if newPerson.CI != personas[newPerson.CI].CI {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": "La CI de una persona no se puede cambiar",
		})
		return
	}

	// personas[newPerson.CI].Birthdate -> Fecha de nacimiento actual de la persona
	if newPerson.Birthdate != personas[newPerson.CI].Birthdate {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": "La fecha de nacimiento de una persona no se puede cambiar",
		})
		return
	}

	personas[newPerson.CI] = newPerson

	actualizarPersonas(personas)
	if err := actualizarPersonas(personas); err != nil {
		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": "Error al actualizar persona de la lista, intentelo mas tarde",
		})
		return
	}

	utils.RJSON(w, http.StatusCreated, utils.JSON{
		"message": fmt.Sprintf("La personas con la CI %d (%s %s) fue modificada correctamente", newPerson.CI, newPerson.Name, newPerson.Surname),
	})
}

// Endpoint - /users/ - DELETE
func deletePerson(w http.ResponseWriter, r *http.Request) {
	persona := utils.People{}
	if err := utils.LJSON(w, r, &persona); err != nil {
		return
	}

	if err := persona.ValidCI(); err != nil {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": err.Error(),
		})
		return
	}

	personas, err := obtenerPersonas()
	if err != nil {
		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": "Error al intentar leer los usuarios, intentelo mas tarde",
		})
		return
	}

	_, ok := personas[persona.CI]
	if !ok {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": fmt.Sprintf("La persona con la CI %d no existe", persona.CI),
		})
		return
	}

	delete(personas, persona.CI)
	if err := actualizarPersonas(personas); err != nil {
		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": "Error al elimiar persona de la lista, intentelo mas tarde",
		})
		return
	}

	utils.RJSON(w, http.StatusOK, utils.JSON{
		"message": fmt.Sprintf("La persona con la cedula %d se ha dado de baja con exito", persona.CI),
	})
}
