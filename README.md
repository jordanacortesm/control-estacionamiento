# control-estacionamiento
 Software open source y gratuito para control de parqueadero y estacionamiento

![Control de parqueadero gratuito](https://parzibyte.me/blog/wp-content/uploads/2021/05/Dashboard-de-sistema-para-estacionamientos-Mostrar-graficas-y-total-de-pagos.png)

# Tutorial, presentación y descargas
En mi blog: https://parzibyte.me/blog/2021/05/11/sistema-estacionamientos-gratuito-open-source/

# Documentación
Las tablas o mejor dicho el esquema de la base de datos se encuentra en **bd.go** dentro de la función
`crearTablas`

# Preparando para producción
Vas a necesitar GCC de 64 bits para compilar SQLite3, y obviamente el compilador de Go.

Primero prepara el frontend con:

`npm run build`

Ahora compila la versión para producción en el backend con:

`make prod`

Crea una nueva carpeta limpia y:

1. Coloca dentro el ejecutable (normalmente es **control_estacionamiento_prod_64.exe**).
2. Igualmente ahí crea la carpeta llamada **public** 
3. Copia lo que está en **frontend/dist** excluyendo los archivos **.map** dentro de **public**

Al final debe quedarte algo así:
```
Listado de rutas de carpetas
El número de serie del volumen es 
C:.
│   control_estacionamiento_prod_64.exe
│
└───public
    │   favicon.ico
    │   index.html
    │
    ├───css
    │       chunk-vendors.6c494cd6.css
    │
    ├───fonts
    │       materialdesignicons-webfont.147e3378.woff
    │       materialdesignicons-webfont.174c02fc.ttf
    │       materialdesignicons-webfont.64d4cf64.eot
    │       materialdesignicons-webfont.7a44ea19.woff2
    │
    ├───img
    │       parzibyte.55e5dd2e.png
    │
    └───js
            app.e2f01bf2.js
            chunk-vendors.5ccce321.js
```

Luego de eso ya puedes distribuir el software.

# Usando software

Ejecuta el archivo .exe y luego visita `http://localhost:5000/static/` (Mira bien la ruta, escríbela tal cual)
