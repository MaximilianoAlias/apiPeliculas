package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Pelicula struct {
	ID       string    `json:"id"`
	NumPeli  string    `json:"numPeli"`
	Titulo   string    `json:"titulo"`
	Director *Director `json:"director"`
}

type Director struct {
	NombreDir   string `json:"nombre"`
	ApellidoDir string `json:"apellido"`
}

var pelicula []Pelicula

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pelicula)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	/*
		en este caso no necesito usar el index por lo que pongo un guion bajo
		para ignorar esa variable ya que si la declaro y no la uso, el recolector
		de basura de golang me va a dar error.
	*/
	for _, items := range pelicula {
		if items.ID == params["id"] {
			json.NewEncoder(w).Encode(items)
			return
		}
	}

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nuevaPelicula Pelicula
	_ = json.NewDecoder(r.Body).Decode(&pelicula)
	nuevaPelicula.ID = strconv.Itoa(rand.Intn(999))

	pelicula = append(pelicula, nuevaPelicula)

	json.NewEncoder(w).Encode(nuevaPelicula)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, items := range pelicula {
		if items.ID == params["id"] {
			pelicula = append(pelicula[:index], pelicula[index+1:]...)

			var nuevaPelicula Pelicula
			_ = json.NewDecoder(r.Body).Decode(&pelicula)
			nuevaPelicula.ID = strconv.Itoa(rand.Intn(999))

			pelicula = append(pelicula, nuevaPelicula)

			json.NewEncoder(w).Encode(nuevaPelicula)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Buscar la película con el ID dado y eliminarla del slice
	for index, item := range pelicula {
		if item.ID == params["id"] {
			pelicula = append(pelicula[:index], pelicula[index+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Película eliminada correctamente", "deleted_movie": item.Titulo})
			return
		}
	}

	// Si no se encuentra ninguna película con el ID dado, devolver un mensaje de error
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "No se encontró ninguna película con el ID proporcionado"})
}

func main() {
	//crear enrutador
	rutas := mux.NewRouter()

	//agregar peliculas

	pelicula = append(pelicula, Pelicula{ID: "1", NumPeli: "142536", Titulo: "El Rey Leon", Director: &Director{NombreDir: "Maximiliano", ApellidoDir: "Alias"}})
	pelicula = append(pelicula, Pelicula{ID: "2", NumPeli: "789456", Titulo: "La La Land", Director: &Director{NombreDir: "Damien", ApellidoDir: "Chazelle"}})
	pelicula = append(pelicula, Pelicula{ID: "3", NumPeli: "369852", Titulo: "Forrest Gump", Director: &Director{NombreDir: "Robert", ApellidoDir: "Zemeckis"}})
	pelicula = append(pelicula, Pelicula{ID: "4", NumPeli: "951753", Titulo: "Titanic", Director: &Director{NombreDir: "James", ApellidoDir: "Cameron"}})
	pelicula = append(pelicula, Pelicula{ID: "5", NumPeli: "258147", Titulo: "Matrix", Director: &Director{NombreDir: "Lana", ApellidoDir: "Wachowski"}})

	/*
		con HandleFunc de gorilla mux voy a definir la ruta
		que va a ponerse en el navegador separado por una coma,
		seguidamente va colocado el nombre de la funcion que va a
		ejecutar esa ruta "getMovies" y despues inserto el metodo
		que sea acorde a la funcion, como en este caso deseo obtener
		datos, el metodo correspondiente va a ser el "GET".
	*/
	rutas.HandleFunc("/peliculas", getMovies).Methods("GET")
	rutas.HandleFunc("/peliculas/{id}", getMovie).Methods("GET")
	rutas.HandleFunc("/peliculas/nueva", createMovie).Methods("POST")
	rutas.HandleFunc("/peliculas/editar/{id}", updateMovie).Methods("PUT")
	rutas.HandleFunc("peliculas/eliminar/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Iniciando el servidor en el puerto :8000\n")

	//levanto el servidor en el puerto 8000 de mi localhost
	log.Fatal(http.ListenAndServe(":8000", rutas))

}
