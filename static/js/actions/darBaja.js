import { jsonDelete } from "../peticiones.js";
import { exitoMensaje, errorMensaje } from "./mensajes.js";

// Variables con la IDs de los mensajes y del Formulario
const idForm = "#form-baja";
const idMensaje = "#mensaje-baja";

// Envía una petición DELETE al servidor con la cédula del usuario que debe borrar
export function darBajaPersona(e) {
    e.stopPropagation();
    e.preventDefault();

    // Convierto la información del formulario a JSON
    const data = new FormData(document.querySelector(idForm));
    const persona = Object.fromEntries(data.entries());

    // Paso la cédula a Int
    persona.ci = parseInt(persona.ci);

    // Realizo la petición
    jsonDelete("/users/", persona)
        .then((json) => {
            exitoMensaje(idMensaje, json.message);
        })
        .catch((err) => {
            errorMensaje(idMensaje, err.message);
        });
}

export function limpiarCamposBaja(e) {
    e.stopPropagation();
    e.preventDefault();

    // Pregunto si quiere borrar los campos
    if (confirm("¿Desea borrar los campos del \"Baja de Persona\"?")) {
        document.querySelector(idForm).reset();
    }
}
