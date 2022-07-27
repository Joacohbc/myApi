# Documentación de la API

## Endpoint /api

En este endpoints solo se pueden enviar 2 tipos de peticiones: GET y POST.

### GET - /api

La respuesta para una petición GET sera un simple JSON:

```json
{
    "message": "Hola Mundo!"
}
```

### POST - /api

La petición POST debe contener en el body un JSON con un la propiedad "message":

Ejemplo de JSON de la petición:

Petición POST:

```json
{
    "message": "Hola Servidor!"
}
```

Respuesta del servidor:

```json
{
    "message": "El mensaje que enviaste: Hola Servidor!"
}
```

### Cualquier otro tipo de petición

Cualquier petición que no sea GET o POST, recibiría como respuesta un JSON con un mensaje de error

Ejemplo de petición PUT:

```json
{
    "error": "Las peticiones a esta ruta deben ser GET o POST, no se permite: PUT"
}
```

## Endpoint /users/ (Pendiente a terminar)

En este endpoint se aceptan bastantes más tipos de peticiones: GET/HEAD, POST, PUT/PATH y DELETE.

Cada persona tiene los siguientes campos:

- CI, es un campo obligatorio de tipo entero de 8 caracteres. En el JSON el campo es "ci"
- Nombre, es un campo obligatorio de tipo string de mínimo 1 carácter y máximo 50. En el JSON el campo es "name"
- Segundo nombre, es un campo no obligatorio de tipo string de mínimo 1 carácter y máximo 50. En el JSON el campo es "second_name"
- Apellido, es un campo obligatorio de tipo string de mínimo 1 carácter y máximo 50. En el JSON el campo es "surname"
- Segundo apellido, es un campo no obligatorio de tipo string de mínimo 1 carácter y máximo 50. En el JSON el campo es "second_surname"
- Fecha de nacimiento, es un campo obligatorio de tipo que debe ir en el formato "dd/mm/yyyy", usando los "/". En el JSON el campo es "birthdate"

### GET - /users/

Las peticiones GET dirigidas a este endpoint pueden contener, o no, la cédula de la persona de la cual se quiere los datos en URL("/users/12345678").
