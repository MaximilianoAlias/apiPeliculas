consignas:

crear 2 estructuras: Pelicula y Director

la estructura de la pelicula tendra:
id
numeroDePeli
titulo
Director <- va a ser un puntero porque va a estar
            relacionado a la estructura de Director

la estructura de Director:
nombre
apellido

creamos un array llamado peliculas de tipo Pelicula

creamos funcion main

dentro de la main creamos un "enrutador" que la biblioteca
de gorilla mux nos facilita.
ej: r := mux.NewRouter()

seguidamente utilizaremos este enrutador para poder acceder
a las distintas rutas que van a ejecutar las funciones
ej: r.HandleFunc("/pelicula", getMovie).Methods("GET")

creadas todas las rutas que nos hagan falta o creamos necesarias
debemos crear las funciones que aun no hemos creado como por ejemplo:

func getMovie(w http.ResponseWriter, r* http.Request){
    
}