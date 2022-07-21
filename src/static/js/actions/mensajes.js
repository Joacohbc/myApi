// Oculta el mensaje que este mostrando el componente
export function ocultarMensaje(e) {
    e.preventDefault();
    e.stopPropagation();

    e.target.innerHTML = "";
    e.target.setAttribute("class", "");
}

// Muestra un mensaje de éxito dentro del componente (innerHTML) que se le asigne
// el mensaje se oculta luego de 10s
export function exitoMensaje(id,mensaje) {

    let div = document.querySelector(id);

    div.setAttribute("class", "exito");
    div.innerHTML = mensaje;

    setTimeout(() => {
        div.innerHTML = "";
        div.setAttribute("class", "");
    }, 10000);
}

// Muestra un mensaje de error dentro del componente (innerHTML) que se le asigne
// el mensaje se oculta luego de 10s
export function errorMensaje(id,mensaje) {
    let msg = document.querySelector(id);
    
    msg.setAttribute("class", "error");
    msg.innerHTML = mensaje;

    setTimeout(() => {
        msg.innerHTML = "";
        msg.setAttribute("class", "");
    }, 10000);
}
