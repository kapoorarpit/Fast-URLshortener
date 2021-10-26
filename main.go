package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/shorten", CreateEndpoint).Methods("POST")
	router.HandleFunc("/expand/", ExpandEndpoint).Methods("POST")
	router.HandleFunc("/{id}", RedirectEndpoint).Methods("GET")
	router.HandleFunc("/", Home).Methods("GET")
	log.Fatal((http.ListenAndServe(":8000", router)))
}

type URL struct {
	LongURL  string    `json:"longURL,omitempty"`
	shortURL string    `json:"shortURL,omitempty"`
	date     time.Time `json:"date,omitempty"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("<h1>This is the homepage<h1>")))
}

func CreateEndpoint(w http.ResponseWriter, r *http.Request) {

}

func ExpandEndpoint(w http.ResponseWriter, r *http.Request) {

}

func RedirectEndpoint(w http.ResponseWriter, r *http.Request) {

}
