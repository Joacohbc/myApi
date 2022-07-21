// Retorna una petición FETCH GET, sin headers
export async function jsonGet(url = "") {
    try {
        let respuesta = await fetch(url);
        let json = await respuesta.json();

        if (!respuesta.ok) {
            throw new Error(json.error);
        }

        return json;
    } catch (e) {
        throw new Error(e);
    }
}

// Retorna una petición FETCH POST, con body en JSON
export async function jsonPost(url = "", data = {}) {
    // Indico los headers
    let headers = {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };

    try {
        // Realizo la petición al URL indicado con los headers adecuados`
        let respuesta = await fetch(url, headers);

        // Paso la respuesta a JSON
        let json = await respuesta.json();

        // En la petición no tiene un status OK que retorne el error
        if (!respuesta.ok) {
            throw new Error(json.error);
        }

        return json;
    } catch (e) {
        throw new Error(e);
    }
}

// Retorna una petición FETCH DELETE, con body en JSON
export async function jsonDelete(url = "", data = {}) {
    let headers = {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };

    try {
        let respuesta = await fetch(url, headers);
        let json = await respuesta.json();

        if (!respuesta.ok) {
            throw new Error(json.error);
        }

        return json;
    } catch (e) {
        throw new Error(e);
    }
}

// Retorna una petición FETCH PUT, con body en JSON
export async function jsonPut(url = "", data = {}) {
    let headers = {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };

    try {
        let respuesta = await fetch(url, headers);
        let json = await respuesta.json();

        if (!respuesta.ok) {
            throw new Error(json.error);
        }

        return json;
    } catch (e) {
        throw new Error(e);
    }
}
