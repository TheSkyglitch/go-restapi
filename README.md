# PROYECTO REST API
# Daniel Flores de la Torre - 200614

## Uso
1. Se ejecuta el programa main.go
2. Usar Postman para enviar solicitudes a la base de datos
2. Las rutas para la API son:
    * **GET /users:** Esta ruta lista todos los usuarios disponibles en la base de datos.
    * **POST /users:** Crea un nuevo usuario en la base de datos con los atributos que se hayan especificado en el Body del JSON.
    * **PUT /users/{id}:** Con el ID del usuario, podemos actualizar sus datos nuevos en una nueva solicitud JSON. Si no se encuentra el ususario, dará error.
    * **DELETE /users/{id}:** Elimina un usuario utilizando su ID. Si no se encuentra el usuario, dará error.