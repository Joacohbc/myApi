package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"myAPI/src/models"
	"myAPI/src/utils"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const personasJsonPath string = "./personas.json"

func init() {

	// Si el archivo existe de las personas, retorno para continuar con el código
	if _, err := os.Stat(personasJsonPath); err == nil {
		return
	}

	// Sino existe el archivo, creo un nuevo JSON vació
	b, err := json.MarshalIndent(map[int]models.People{}, " ", "\t")
	if err != nil {
		log.Println("Error al crear el archivo de persona:", err)
	}

	// Y creo el archivo
	err = os.WriteFile(personasJsonPath, b, 0644)
	if err != nil {
		log.Println("Error al crear el archivo de persona:", err)
	}
}

// Retorna un Map de personas y una Flag que indica si se realizo con éxito la lectura de personas (true)
// o si ocurrió un error (false)
func obtenerPersonas(w http.ResponseWriter) (map[int]models.People, bool) {

	// Leo el archivo
	b, err := os.ReadFile(personasJsonPath)
	if err != nil {
		log.Println("Error al leer el archivo de personas:", err)

		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": "Error al intentar leer las personas, intentelo mas tarde",
		})
		return nil, false
	}

	// Lo cargo en un map de persona
	var personas map[int]models.People
	if err = json.Unmarshal(b, &personas); err != nil {
		log.Println("Error al cargar el archivo de personas al map de personas:", err)

		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": "Error al intentar leer las personas, intentelo mas tarde",
		})
		return nil, false
	}

	// Y lo retorno
	return personas, true
}

// Sobrescribe el archivo json de las personas
func actualizarPersonas(persona map[int]models.People) error {
	b, err := json.MarshalIndent(persona, " ", "\t")
	if err != nil {
		log.Println("Error al actualizar el archivo de personas:", err)
		return err
	}

	err = os.WriteFile(personasJsonPath, b, 0644)
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

	// Leo el URL en busca de la CI de la persona que me piden
	ci, err := strconv.Atoi(utils.GetLastPathVariable(r, "/users/"))
	if err != nil {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": "la cedula persona pedida es invalida",
		})
		return
	}

	// Creo una persona y le asigno esa cédula
	personaPedida := models.People{
		CI: ci,
	}

	// Valido que esa cédula sea correcta
	if err := personaPedida.ValidCI(); err != nil {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": err.Error(),
		})
		return
	}

	// Leo todas las personas
	personas, exito := obtenerPersonas(w)
	if !exito {
		return
	}

	// Si la persona no existe en el map, lo informo con un error
	_, ok := personas[personaPedida.CI]
	if !ok {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": fmt.Sprintf("La persona con la CI %d no existe", personaPedida.CI),
		})
		return
	}

	// Si existe la retorno
	utils.RJSON(w, http.StatusOK, personas[personaPedida.CI])
}

// Endpoint - /users/ - GET/HEAD
func getAllPeople(w http.ResponseWriter, _ *http.Request) {

	// Obtengo el listado de todas las personas
	personasMap, exito := obtenerPersonas(w)
	if !exito {
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

// Endpoint - /users/ - POST
func newPerson(w http.ResponseWriter, r *http.Request) {

	// Leo la persona del body de la petición
	persona := models.People{}
	if err := utils.LJSON(w, r, &persona); err != nil {
		return
	}

	// Valido que tenga todos los campos correctos
	if err := persona.ValidAll(); err != nil {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": err.Error(),
		})
		return
	}

	// Obtengo el map de personas
	personas, exito := obtenerPersonas(w)
	if !exito {
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

	// Obtengo la CI de la persona que quiero modificar
	ci, err := strconv.Atoi(utils.GetLastPathVariable(r, "/users/"))
	if err != nil {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": "la cedula persona pedida es invalida",
		})
		return
	}

	// Obtengo al nueva persona del body de la petición
	newPerson := models.People{}
	if err := utils.LJSON(w, r, &newPerson); err != nil {
		return
	}

	// Le asigno la cedula que se pidio al nuevo usuario para poder validarla
	newPerson.CI = ci
	if err := newPerson.ValidCI(); err != nil {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": err.Error(),
		})
		return
	}

	// Obtengo el map de personas
	personas, exito := obtenerPersonas(w)
	if !exito {
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
		oldPerson := personas[newPerson.CI]

		// Verifico que la cédula no se haya cambiado (por las dudas)
		if newPerson.CI != oldPerson.CI {
			utils.RJSON(w, http.StatusBadRequest, utils.JSON{
				"error": "La CI de una persona no se puede cambiar",
			})
			return
		}

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
		if err := newPerson.ValidAll(); err != nil {
			utils.RJSON(w, http.StatusBadRequest, utils.JSON{
				"error": err.Error(),
			})
			return
		}
	}

	// Sobrescribo la persona y actualizo el archivo
	personas[newPerson.CI] = newPerson
	if err := actualizarPersonas(personas); err != nil {
		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": "Error al actualizar persona de la lista, intentelo mas tarde",
		})
		return
	}

	// Envió mensaje de éxito
	utils.RJSON(w, http.StatusCreated, utils.JSON{
		"message": fmt.Sprintf("La personas con la CI %d (%s %s) fue modificada correctamente", newPerson.CI, newPerson.Name, newPerson.Surname),
	})
}

// Endpoint - /users/ - DELETE
func deletePerson(w http.ResponseWriter, r *http.Request) {

	// Obtengo la CI de la persona que quiero modificar
	ci, err := strconv.Atoi(utils.GetLastPathVariable(r, "/users/"))
	if err != nil {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": "la cedula persona pedida es invalida",
		})
		return
	}

	persona := models.People{
		CI: ci,
	}

	// Valido la cédula de la persona
	if err := persona.ValidCI(); err != nil {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": err.Error(),
		})
		return
	}

	// Obtengo el map de personas
	personas, exito := obtenerPersonas(w)
	if !exito {
		return
	}

	// Verifico que la persona exista, para poder eliminarla
	_, ok := personas[ci]
	if !ok {
		utils.RJSON(w, http.StatusBadRequest, utils.JSON{
			"error": fmt.Sprintf("La persona con la CI %d no existe", persona.CI),
		})
		return
	}

	// Elimino la persona dle map y actualizo el archivo
	delete(personas, ci)
	if err := actualizarPersonas(personas); err != nil {
		utils.RJSON(w, http.StatusInternalServerError, utils.JSON{
			"error": "Error al elimiar persona de la lista, intentelo mas tarde",
		})
		return
	}

	// Envió mensaje de éxito
	utils.RJSON(w, http.StatusOK, utils.JSON{
		"message": fmt.Sprintf("La persona con la cedula %d se ha dado de baja con exito", persona.CI),
	})
}
