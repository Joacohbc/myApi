import { jsonGet, jsonPut } from "../peticiones.js";
import { errorMensaje, exitoMensaje } from "./mensajes.js";

// Determina la visibilidad del Form.
/*
    En caso de ocultarse, se deja solo el campo de la cédula vació y se le re-asigna la acción al botón de cargar.

    En caso de mostrase se pedirá un dato adicional que es el JSON con los datos de la persona para cargar el 
    formulario completo y ree-asignar las acciones a los botones de modificar y cancelar (y borrar el botón de 
    cargar).
*/
function visibilidadForm(idForm, mostrar, json = {}) {
    // Selecciono el form mostrar/ocultar
    let form = document.querySelector(idForm);

    // Si me piden mostrar el form
    if (mostrar) {
        // Le cargo el HTML con la información dentro
        form.innerHTML = `
        <label> CI: </label> 
        <input type="number" name="ci" value="${json.ci}"readonly> <br>

        <label> Nombres: </label> 
        <input type="text" name="name" placeholder="Primer nombre" value="${json.name}" required>
        <input type="text" name="second_name" placeholder="Segundo nombre" value="${json.second_name}"> <br>

        <label> Apellidos: </label>
        <input type="text" name="surname" placeholder="Primer apellido" value="${json.surname}" required>
        <input type="text" name="second_surname"  placeholder="Segundo apellido" value="${json.second_surname}"> <br>

        <label> Fecha de nacimiento: </label>
        <input type="text" name="birthdate" value="${json.birthdate}" readonly> <br>  

        <input type="submit" id="btnModif" value="Modificación">
        <input type="submit" id="btnCancelarModif" value="Cancelar"> 
        `;

        // Y le asigno los eventos a los botones
        document
            .querySelector("#btnModif")
            .addEventListener("click", modificarPersona);

        document
            .querySelector("#btnCancelarModif")
            .addEventListener("click", cancelarModifiacion);
    } else {
        // Sobrescribo el contenido HTML solo dejando los datos para cargar a la persona
        form.innerHTML = `
        <label> CI: </label> 
        <input type="text" name="ci" id="ci" placeholder="Cédula sin guiones" required> 
        <button id="btnCargarDatos"> Cargar </button> <br>
        `;

        // Y el evento al botón
        document
            .querySelector("#btnCargarDatos")
            .addEventListener("click", cargarPersona);
    }
}

// Envía una petición PUT al servidor con los nuevo dato de la persona para que los modifiquen
function modificarPersona(e) {
    e.preventDefault();
    e.stopPropagation();

    // Variables con la IDs de los mensajes y del Formulario
    const idMensaje = "#mensaje-modif";
    const idForm = "#form-modif";

    // Convierto la información del formulario a JSON
    const data = new FormData(document.querySelector(idForm));
    const persona = Object.fromEntries(data.entries());

    // Paso la cédula a Int
    persona.ci = parseInt(persona.ci);
    if (isNaN(persona.ci)) {
        errorMensaje(idMensaje, "La cedula debe estar vacia");
    }

    // Realizo la petición PUT y le paso los datos de la nueva persona
    jsonPut("/users", persona)
        .then((json) => {
            // Vació y oculto el formulario
            visibilidadForm(idForm, false);

            // Muestro un mensaje de éxito
            exitoMensaje(idMensaje, json.message);
        })
        .catch((err) => {
            // Si ocurrió un error lo muestro
            errorMensaje(idMensaje, err.message);
        });
}

// Oculta el formulario de modificación
function cancelarModifiacion(e) {
    e.preventDefault();
    e.stopPropagation();

    // Variables con la IDs de los mensajes y del Formulario
    const idMensaje = "#mensaje-modif";
    const idForm = "#form-modif";

    // Vació y oculto el formulario
    visibilidadForm(idForm, false);

    // Muestro un mensaje de que se cancelo la modificación
    exitoMensaje(idMensaje, "Modificacion cancelada con exito");
}

// Envía una petición GET al servidor pidiendo los datos y los carga en el formulario de modificación
export function cargarPersona(e) {
    e.preventDefault();
    e.stopPropagation();

    // Variables con la IDs de los mensajes y del Formulario
    const idMensaje = "#mensaje-modif";
    const idForm = "#form-modif";

    // Obtengo los datos del formulario
    const data = new FormData(document.querySelector(idForm));

    // Paso la cédula a Int
    data.set('ci', parseInt(data.get('ci')));
    if (isNaN(data.get('ci'))) {
        errorMensaje(idMensaje, "La cedula debe estar vacia");
    }

    // Envió la petición a /users/{ci} y obtengo los datos
    jsonGet(`/users/${data.get("ci")}`)
        .then((json) => {
            // Muestro el formulario y lo cargo
            visibilidadForm(idForm, true, json);

            // Muestro un mensaje de éxito
            exitoMensaje(idMensaje, "Persona cargada con exito");
        })
        .catch((err) => {
            // Si ocurrió un error lo muestro
            errorMensaje(idMensaje, err.message);
        });
}
