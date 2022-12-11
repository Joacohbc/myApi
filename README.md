# My API REST

Una API REST desarrollada en Go 100% Vanilla de un 3 simples endpoints los cuales permite realizar:

1. Para comprobar estado del servidor. Endpoint: /api  que acepta peticiones GET y POST.
2. Para realizar altas, bajas, modificación y consulta de personas. Endpoint: /users/ que acepta peticiones GET/HEAD, POST, PUT/PATH y DELETE
3. La API sirve un Front-End muy simple desarrollado con HTML/CSS y Javascript 100% Vanilla para comprobar el funcionamiento del API mediante peticiones realizar con FETCH.

Adicionalmente la API guarda los LOGs en en la carpeta logs/ del directorio actual(usar -log=false) para desactivar la opción.

## Abrir servidor

Simplemente hay que ejecutar el archivo binario correspondiente (en este caso, myapi64.bin porque es un Linux x64) para poder iniciar el servidor, si se quiere detener basta con hacer Ctrl+C (^C).

```bash
git clone https://github.com/Joacohbc/myApi; 
cd myApi; 
chmod +x ./bin/myapi64.bin; # Para dar permisos de ejecución
```

```bash
# Las flags -port para indiciar el puerto de escucha y -front para indicar la carpeta donde se ubica los archivos HTML/CSS y JavaScript
./bin/myapi-64.bin -port 8080 -front ./static
__  __          _    ____ ___   ____           _   
|  \/  |_   _   / \  |  _ \_ _| |  _ \ ___  ___| |_ 
| |\/| | | | | / _ \ | |_) | |  | |_) / _ \/ __| __|
| |  | | |_| |/ ___ \|  __/| |  |  _ <  __/\__ \ |_ 
|_|  |_|\__, /_/   \_\_|  |___| |_| \_\___||___/\__|
        |___/ Version: v1.0
- Puerto: 8080
- Frontend: /home/user/Projects/myApi/static
- Inicio: 07/21/2022 23:05:08

Servidor escuchando...
```

## Documentación

Aquí un poco de documentación del API: [DOCUMENTATION.md](./docs/DOCUMENTATION.md)
