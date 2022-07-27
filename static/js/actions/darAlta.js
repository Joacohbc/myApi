import { jsonPost } from "../peticiones.js";
import { exitoMensaje, errorMensaje } from "./mensajes.js";

export function darAltaPersona(e) {
    e.stopPropagation();
    e.preventDefault();

    const idMensaje = "#mensaje-alta";
    const idForm = "#form-alta";

    // Convierto la información del formulario a JSON
    const data = new FormData(document.querySelector(idForm));
    const persona = Object.fromEntries(data.entries());

    // Paso la cédula a Int
    persona.ci = parseInt(persona.ci);

    // Paso la fecha al formato que pide el servidor: yyyy-MM-dd -> dd/MM/yyyy
    let fecha = persona.birthdate.split("-");
    persona.birthdate = fecha[2] + "/" + fecha[1] + "/" + fecha[0];

    // Realizo la petición POST enviado lso datos de la persona
    jsonPost("/users/", persona)
        .then((json) => {

            // Muestra un mensaje de éxito 
            exitoMensaje(idMensaje, json.message);

            // Pregunto si quiere borrar los campos
            if (confirm("¿Desea borrar los campos?")) {
                document.querySelector(idForm).reset();
            }
        })
        .catch((err) => {
            // Si ocurrió un error muestra el error
            errorMensaje(idMensaje, err.message);
        });
}
