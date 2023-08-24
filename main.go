package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"
	// "strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"fristname"`
	Lastname  string `json:"lastname"`
}

func main() {

	movies = append(movies, Movie{Id: "1", Isbn: "1234", Title: "Movie one", Director: &Director{Firstname: "Jithu", Lastname: "joseph"}})
	movies = append(movies, Movie{Id: "2", Isbn: "2345", Title: "Movie two", Director: &Director{Firstname: "vineeth", Lastname: "sreenivasan"}})

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

var movies []Movie

// get movies function

func getMovies(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

// get one movie

func getMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(req)
	for _, item := range movies {
		if item.Id == param["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

// create movie

func createMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie

	json.NewDecoder(req.Body).Decode(&movie)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)

}

// delete movie

func deleteMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(req)
	for idx, item := range movies {
		if item.Id == param["id"] {
			// delete concept
			movies = append(movies[:idx], movies[idx+1:]...)
			break
		}

	}
}

// update movie

func updateMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// first delete with id ..then create
	param := mux.Vars(req)
	var movie Movie
	for idx, item := range movies {
		if item.Id == param["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)

			json.NewDecoder(req.Body).Decode(&movie)
			movie.Id = param["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}
