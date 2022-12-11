# Documentación de la API

Documentacion en POSTMAN: [Postman_myApi_Documentation.json](./myAPI%20Native%20Golang%20Documentation.json)

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

## Endpoint /users/

En este endpoint se aceptan bastantes más tipos de peticiones: GET/HEAD, POST, PUT/PATH y DELETE.

Cada persona tiene los siguientes campos:

- CI, es un campo obligatorio de tipo entero de 8 caracteres. En el JSON el campo es "ci"
- Nombre, es un campo obligatorio de tipo string de mínimo 1 carácter y máximo 50. En el JSON el campo es "name"
- Segundo nombre, es un campo no obligatorio de tipo string de mínimo 1 carácter y máximo 50. En el JSON el campo es "second_name"
- Apellido, es un campo obligatorio de tipo string de mínimo 1 carácter y máximo 50. En el JSON el campo es "surname"
- Segundo apellido, es un campo no obligatorio de tipo string de mínimo 1 carácter y máximo 50. En el JSON el campo es "second_surname"
- Fecha de nacimiento, es un campo obligatorio de tipo que debe ir en el formato "dd/mm/yyyy", usando los "/". En el JSON el campo es "birthdate". Ejemplo 01/01/2001, no sirve la fecha 1/1/2001, debe incluir 2 digitos para el mes y el dia

Adicionalmente cuando ser pide un listado de la información de las personas personas viene con un dato adicional, el campo "birthdate_time" que es un campo que solo utiliza el servidor en el tipo time.Time

### GET/HEAD - /users/

Las peticiones GET/HEAD dirigidas a este endpoint pueden contener, o no, la cédula de la persona de la cual se quiere los datos en URL("/users/58762269") para retornar unicamente la persona indicada, o no contener nada ("/users/") para obtener un listado de todas las personas registradas.

Ejemplo de un petición GET a "/users/58762269":

```json
{
    "ci": 58762269,
    "name": "Tara",
    "second_name": "Lany",
    "surname": "Rosenfeld",
    "second_surname": "Burges",
    "birthdate": "19/01/1964",
    "birthdate_time": "1964-01-19T00:00:00Z",
}
```

Ejemplo de un petición GET a "/users/":

```json
[
    {
        "ci": 58762269,
        "name": "Tara",
        "second_name": "Lany",
        "surname": "Rosenfeld",
        "second_surname": "Burges",
        "birthdate": "19/01/1964",
        "birthdate_time": "1964-01-19T00:00:00Z",
    },
    {
        "ci": 71461179,
        "name": "Joseph",
        "second_name": "",
        "surname": "Karr",
        "second_surname": "Celestine",
        "birthdate": "03/01/1952",
        "birthdate_time": "1952-01-03T00:00:00Z",
    },
    {
        "ci": 12345678,
        "name": "Barbara",
        "second_name": "Liliana",
        "surname": "Eldredg",
        "second_surname": "",
        "birthdate": "15/03/1983",
        "birthdate_time": "1983-03-15T00:00:00Z",
    }
]
```

### POST - /users/

Las peticiones POST dirigidas a este endpoint deben contener un JSON en el body de la petición contenga como mínimo todos los campos obligatorios. La cédula al ser identificador de cada persona no se puede repetir, y en caso de que se repita el servidor informara del error. Los campos no obligatorios no tiene porque estar dentro del JSON (en caso de que no se quieran ingresar).

Ejemplo de una petición POST con un JSON en el body:

```json
{
    "ci": 57960390,
    "name": "Pete",
    "second_name": "",
    "surname": "Little",
    "second_surname": "",
    "birthdate": "12/02/1984"
}
```

Siendo lo mismos que:

```json
{
    "ci": 57960390,
    "name": "Pete",
    "surname": "Little",
    "birthdate": "12/02/1984"
}
```

Ejemplo de respuesta del servidor:

```json
{
    "message": "La persona Pete Little fue creada con éxito"
}
```

### DELETE - /users/

Las peticiones DELETE dirigidas a este endpoint deben contener una CI de la persona que se quiere eliminar en la URL de la petición.

Ejemplo de petición DELETE a /users/12345678

Ejemplo de respuesta del servidor:

```json
{
    "message": "La persona con la cédula 12345678 se ha dado de baja con éxito"
}
```

### PATCH/PUT - /users/

Las peticiones PATCH/PUT dirigidas a este endpoint deben contener un JSON en el body de la petición contenga todos los campos de la persona que se quiere actualizar, no se pueden actualizar la fecha de nacimiento ni la cédula de la persona. **Importante**: Si se dejan en blanco los campos o se omiten estos se tomaran como "no modificados" y mantendrán su valor actual

Persona original:

```json
{
    "ci": 58762269,
    "name": "Tara",
    "second_name": "Lany",
    "surname": "Rosenfeld",
    "second_surname": "Burges",
    "birthdate": "19/01/1964"
}
```

Ejemplo de petición PUT a /users/58762269:

```json
{
    "name": "Moly",
    "second_name": "Daniela",
    "surname": "Rosenfeld",
    "second_surname": "",
    "birthdate": "19/01/1964"
}
```

Ejemplo de respuesta del servidor:

```json
{
    "message": "La personas con la CI 58762269 (Moly Rosenfeld) fue modificada correctamente"
}
```
