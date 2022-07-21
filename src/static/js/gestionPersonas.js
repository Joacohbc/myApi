import { darAltaPersona } from "./actions/darAlta.js";
import { darBajaPersona } from "./actions/darBaja.js";
import { cargarPersona} from "./actions/modificarPersona.js";
import { ocultarMensaje } from "./actions/mensajes.js";

document
    .querySelector("#mensaje-alta")
    .addEventListener("click", ocultarMensaje);

document
    .querySelector("#mensaje-baja")
    .addEventListener("click", ocultarMensaje);

document
    .querySelector("#mensaje-modif")
    .addEventListener("click", ocultarMensaje);

document.querySelector("#btnAlta").addEventListener("click", darAltaPersona);

document.querySelector("#btnBaja").addEventListener("click", darBajaPersona);

document
    .querySelector("#btnCargarDatos")
    .addEventListener("click", cargarPersona);
