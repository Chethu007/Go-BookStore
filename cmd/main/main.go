package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/Chethu007/Go-Code/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func RegisterMoviesRoute(r2 *mux.Router) {
	r2.HandleFunc("/movie/", CreateMovie).Methods("POST")
	r2.HandleFunc("/movie/", GetAllMovies).Methods("GET")
	r2.HandleFunc("/movie/{id}", GetMovieById).Methods("GET")
	r2.HandleFunc("/movie/{id}", UpdateMovie).Methods("PUT")
	r2.HandleFunc("/movie/{id}", DeleteMovie).Methods("DELETE")
}

type Movie struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Director Director `json:"director"`
	Year     string   `json:"year"`
}

type Director struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var movies []Movie

func main() {
	movies = append(movies, Movie{"1", "Forest Gump", Director{"Robert", 70}, "1995"})
	movies = append(movies, Movie{"2", "Seven", Director{"David Fincher", 60}, "1998"})

	fmt.Println("CRUD mini project started........")
	fmt.Println("Book Route Registered........")
	r := mux.NewRouter()
	routes.RegisterBookStoreRouted(r)

	//without db
	fmt.Println("Movie Route Registered........")
	r2 := mux.NewRouter()
	RegisterMoviesRoute(r2)

	go func() {
		fmt.Println("Server 1 listening on port 9010")
		log.Fatal(http.ListenAndServe("localhost:9010", r))
	}()

	go func() {
		fmt.Println("Server 1 listening on port 9090")
		log.Fatal(http.ListenAndServe("localhost:9090", r2))
	}()

	select {}

}

// Handler Func
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		fmt.Println("Error while deccoding json body(Create)")
	}
	if movie.ID == "" {
		movie.ID = strconv.Itoa(rand.Intn(1000000))
	}
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for _, item := range movies {
		if item.ID == param["id"] { //x, _ := strconv.Atoi(param["id"]); item.ID == x {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		fmt.Println("Error while deccoding json body(Update)")
	}
	params := mux.Vars(r)
	for i, item := range movies {
		if item.ID == params["id"] { //x, _ := strconv.Atoi(params["id"]); item.ID == x {
			movies = append(movies[:i], movies[i+1:]...)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range movies {
		if item.ID == params["id"] { //x, _ := strconv.Atoi(params["id"]); item.ID == x {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
