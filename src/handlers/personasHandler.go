package handlers

import (
	"fmt"
	"myAPI/src/logger"
	"myAPI/src/models"
	"myAPI/src/utils"
	"net/http"
	"strconv"
)

const (

	// URL donde se llama a este Endpoint
	URLServed string = "/users/"
)

// Endpoint - /users/ - Gestiona todas las peticiones
func Personas(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET", "HEAD":
		if utils.GetLastPathVariable(r, URLServed) == "" {
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
		logger.Logger.Println("Se hizo una petición:", r.Method)
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": "Las peticiones a esta ruta deben ser GET, HEAD, POST, DELETE, PUT o PATCH",
		})
	}
}

// Endpoint - GET/HEAD
func getPerson(w http.ResponseWriter, r *http.Request) {

	// Leo el URL en busca de la CI de la persona que me piden
	ci, err := strconv.Atoi(utils.GetLastPathVariable(r, URLServed))
	if err != nil {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": "la cédula pedida es invalida",
		})
		return
	}

	// Valido que esa cédula sea correcta
	if err := models.ValidCI(ci); err != nil {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": err.Error(),
		})
		return
	}

	// Leo todas las personas
	personas, err := models.PeopleService().ObtenerPersonas(w)
	if err != nil {
		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": err.Error(),
		})
		return
	}

	// Si la persona no existe en el map, lo informo con un error
	per, ok := personas[ci]
	if !ok {
		utils.RJSON(w, http.StatusNotFound, utils.JSON{
			"error": fmt.Sprintf("La persona con la CI %d no existe", ci),
		})
		return
	}

	// Si existe la retorno
	utils.RJSON(w, http.StatusOK, per)
}

// Endpoint - GET/HEAD
func getAllPeople(w http.ResponseWriter, _ *http.Request) {

	// Obtengo el listado de todas las personas
	personasMap, err := models.PeopleService().ObtenerPersonas(w)
	if err != nil {
		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": err.Error(),
		})
		return
	}

	// Guardo los datos del map en una slice para enviarlos
	personas := []models.People{}
	for _, v := range personasMap {
		personas = append(personas, v)
	}

	// Y envió el slice de personas
	utils.RJSON(w, http.StatusOK, personas)
}

// Endpoint - POST
func newPerson(w http.ResponseWriter, r *http.Request) {

	// Leo la persona del body de la petición
	persona := models.People{}
	if err := utils.LJSON(w, r, &persona); err != nil {
		return
	}

	// Valido que tenga todos los campos correctos
	if err := persona.FormatAndValidAll(); err != nil {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": err.Error(),
		})
		return
	}

	// Obtengo el map de personas
	personas, err := models.PeopleService().ObtenerPersonas(w)
	if err != nil {
		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": err.Error(),
		})
		return
	}

	// Verifico que la persona no exista ya en el listado
	_, ok := personas[persona.CI]
	if ok {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": fmt.Sprintf("La persona con la CI %d ya existe", persona.CI),
		})
		return
	}

	// Agrego a la persona y actualizo el archivo
	personas[persona.CI] = persona
	if err := models.PeopleService().ActualizarPersonas(personas); err != nil {
		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": "Error al agregar a la nueva persona a la lista, inténtelo mas tarde",
		})
		return
	}

	utils.RJSON(w, http.StatusCreated, utils.JSON{
		"message": fmt.Sprintf("La persona %s %s fue creada con éxito", persona.Name, persona.Surname),
	})
}

// Endpoint - PUT/PATCH
func updatePerson(w http.ResponseWriter, r *http.Request) {

	// Obtengo la CI de la persona que quiero modificar
	ci, err := strconv.Atoi(utils.GetLastPathVariable(r, URLServed))
	if err != nil {
		utils.RJSON(w, http.StatusNotFound, utils.JSON{
			"error": "la cédula pedida es invalida",
		})
		return
	}

	if err := models.ValidCI(ci); err != nil {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": err.Error(),
		})
		return
	}

	// Obtengo al nueva persona del body de la petición
	newPerson := models.People{}
	if err := utils.LJSON(w, r, &newPerson); err != nil {
		return
	}

	// Obtengo el map de personas
	personas, err := models.PeopleService().ObtenerPersonas(w)
	if err != nil {
		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": err.Error(),
		})
		return
	}

	// Verifico que la persona exista para poder modificarla
	_, ok := personas[ci]
	if !ok {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": fmt.Sprintf("La persona con la CI %d no existe", newPerson.CI),
		})
		return
	}

	// Si hay algún campo esta vació (es decir que no se va modificar) lo sobrescribo con su valor actual (oldPerson)
	{

		// Guardo los valores actuales de la persona en oldPerson
		oldPerson := personas[ci]

		// Le asigno la cédula a la persona (Para luego en al validación no tire erro por no tenerla)
		newPerson.CI = ci

		//
		// Si algún campo esta vació, significa que no se quiere modificar
		//
		if newPerson.Name == "" {
			newPerson.Name = oldPerson.Name
		}

		if newPerson.SecondName == "" {
			newPerson.SecondName = oldPerson.SecondName
		}

		if newPerson.Surname == "" {
			newPerson.Surname = oldPerson.Surname
		}

		if newPerson.SecondSurname == "" {
			newPerson.SecondSurname = oldPerson.SecondSurname
		}

		if newPerson.Birthdate == "" {
			newPerson.Birthdate = oldPerson.Birthdate
		}

		if newPerson.Birthdate != oldPerson.Birthdate {
			utils.RJSON(w, http.StatusBadRequest, utils.JSON{
				"error": "La fecha de nacimiento de una persona no se puede cambiar",
			})
			return
		}

		//Valido todos los campos de la persona
		if err := newPerson.FormatAndValidAll(); err != nil {
			utils.RJSON(w, http.StatusBadRequest, utils.JSON{
				"error": err.Error(),
			})
			return
		}
	}

	// Sobrescribo la persona y actualizo el archivo
	personas[ci] = newPerson
	if err := models.PeopleService().ActualizarPersonas(personas); err != nil {
		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": "Error al actualizar persona de la lista, inténtelo mas tarde",
		})
		return
	}

	// Envió mensaje de éxito
	utils.RJSON(w, http.StatusCreated, utils.JSON{
		"message": fmt.Sprintf("La personas con la CI %d (%s %s) fue modificada correctamente", newPerson.CI, newPerson.Name, newPerson.Surname),
	})
}

// Endpoint - DELETE
func deletePerson(w http.ResponseWriter, r *http.Request) {

	// Obtengo la CI de la persona que quiero modificar
	ci, err := strconv.Atoi(utils.GetLastPathVariable(r, URLServed))
	if err != nil {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": "la cédula persona pedida es invalida",
		})
		return
	}

	// Valido la cédula de la persona
	if err := models.ValidCI(ci); err != nil {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": err.Error(),
		})
		return
	}

	// Obtengo el map de personas
	personas, err := models.PeopleService().ObtenerPersonas(w)
	if err != nil {
		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": err.Error(),
		})
		return
	}

	// Verifico que la persona exista, para poder eliminarla
	_, ok := personas[ci]
	if !ok {
		utils.RJSON(w, http.StatusNotFound, utils.JSON{
			"error": fmt.Sprintf("La persona con la CI %d no existe", ci),
		})
		return
	}

	// Elimino la persona dle map y actualizo el archivo
	delete(personas, ci)
	if err := models.PeopleService().ActualizarPersonas(personas); err != nil {
		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": "Error al eliminar persona de la lista, inténtelo mas tarde",
		})
		return
	}

	// Envió mensaje de éxito
	utils.RJSON(w, http.StatusOK, utils.JSON{
		"message": fmt.Sprintf("La persona con la cédula %d se ha dado de baja con éxito", ci),
	})
}
