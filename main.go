package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type movie struct {
	ID       string     `json:"ID"`
	Isbn     string     `json:"isbn"`
	Title    string     `json:"title"`
	Director *directors `json:"director"`
}

type directors struct {
	First_name string `json:"first name"`
	Last_name  string `json:"last name"`
}

var movies []movie

func ListAllMovies(w http.ResponseWriter, r *http.Request) {
	// list all movies
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func CreateMovies(w http.ResponseWriter, r *http.Request) {
	// create an entrie
	w.Header().Set("Content-Type", "application/json")
	var movie movie

	json.NewDecoder(r.Body).Decode(&movie)

	lenOfList := len(movies)
	movieId := lenOfList - 1
	movie.ID = strconv.Itoa(movieId + 2)
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)

}

func MovieDetails(w http.ResponseWriter, r *http.Request) {
	// function which servers 3 methods ["GET","PUT","DELETE"]
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		param := mux.Vars(r)

		for _, items := range movies {
			if items.ID == param["id"] {
				json.NewEncoder(w).Encode(items)
				return

			}
		}
	}

	if r.Method == "PUT" {
		params := mux.Vars(r)
		for index, items := range movies {
			if items.ID == params["id"] {
				movies = append(movies[:index], movies[index+1:]...)
				var movie movie
				json.NewDecoder(r.Body).Decode(&movie)
				movie.ID = params["id"]
				movies = append(movies, movie)
				json.NewEncoder(w).Encode(movie)
				return

			}
		}
		json.NewEncoder(w).Encode(movies)
	}
	if r.Method == "DELETE" {
		params := mux.Vars(r)

		for index, items := range movies {
			if items.ID == params["id"] {
				movies = append(movies[:index], movies[index+1:]...)
				break
			}
		}
	}

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, movie{ID: "1", Isbn: "4518158s58", Title: "One piece red", Director: &directors{First_name: "Eiichiro", Last_name: "Oda"}})

	movies = append(movies, movie{ID: "2s", Isbn: "451856s58", Title: "Dune 1", Director: &directors{First_name: "John", Last_name: "Doe"}})

	r.HandleFunc("/movies", ListAllMovies).Methods("GET") // get all movies
	r.HandleFunc("/movies", CreateMovies).Methods("POST") // add movies

	r.HandleFunc("/movies/{id}", MovieDetails).Methods("GET")
	r.HandleFunc("/movies/{id}", MovieDetails).Methods("PUT")
	r.HandleFunc("/movies/{id}", MovieDetails).Methods("DELETE")

	fmt.Println("server is starting at 8000")

	// err := http.ListenAndServe("localhost:8000", r)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	log.Fatal(http.ListenAndServe(":8000", r))

}
