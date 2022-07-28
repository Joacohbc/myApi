import { darAltaPersona, limpiarCamposAlta } from "./actions/darAlta.js";
import { darBajaPersona, limpiarCamposBaja } from "./actions/darBaja.js";
import { cargarPersona } from "./actions/modificarPersona.js";
import { ocultarMensaje } from "./actions/mensajes.js";

// Componentes de Alta
document.querySelector("#btnAlta").addEventListener("click", darAltaPersona);

document
    .querySelector("#mensaje-alta")
    .addEventListener("click", ocultarMensaje);

document
    .querySelector("#btnLimpiarAlta")
    .addEventListener("click", limpiarCamposAlta);

// Componentes de Baja
document
    .querySelector("#mensaje-baja")
    .addEventListener("click", ocultarMensaje);

document.querySelector("#btnBaja").addEventListener("click", darBajaPersona);

document
    .querySelector("#btnLimpiarBaja")
    .addEventListener("click", limpiarCamposBaja);

// Componentes de modificación
document
    .querySelector("#mensaje-modif")
    .addEventListener("click", ocultarMensaje);

document
    .querySelector("#btnCargarDatos")
    .addEventListener("click", cargarPersona);
