import { errorMensaje } from "./actions/mensajes.js";
import { jsonGet } from "./peticiones.js";

let tabla = document.querySelector("#personas");

const idMensaje = "#mensaje-consultas";

const unaSolaFila = (msg) => {
    let row = document.createElement("tr");
    let data = document.createElement("td");
    data.innerHTML = msg;
    data.setAttribute("colspan", "6");
    data.setAttribute("text-align", "center");

    row.appendChild(data);
    tabla.appendChild(row);
};

jsonGet("/users/")
    .then((json) => {
        let fragment = document.createDocumentFragment();

        if (json.length == 0) {
            unaSolaFila("No hay personas listadas");
            return;
        }

        console.log(json);

        const arrayJson = Array.from(json.values());

        arrayJson.forEach((persona) => {
            let row = document.createElement("tr");

            let ci = document.createElement("td");
            ci.innerHTML = persona.ci;
            row.appendChild(ci);

            let nombre = document.createElement("td");
            nombre.innerHTML = persona.name;
            row.appendChild(nombre);

            let segNom = document.createElement("td");
            segNom.innerHTML = persona.second_name;
            row.appendChild(segNom);

            let apellido = document.createElement("td");
            apellido.innerHTML = persona.surname;
            row.appendChild(apellido);

            let segApe = document.createElement("td");
            segApe.innerHTML = persona.second_surname;
            row.appendChild(segApe);

            let birthdate = document.createElement("td");
            birthdate.innerHTML = persona.birthdate;
            row.appendChild(birthdate);

            fragment.appendChild(row);
        });

        tabla.appendChild(fragment);
    })
    .catch((err) => {
        unaSolaFila("Error al cargar las personas desde el servidor");

        errorMensaje(idMensaje, err.message);
    });
