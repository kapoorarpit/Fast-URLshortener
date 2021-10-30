package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//ctx, _ := context.WithTimeout(context.Background(), 10+time.Second)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection = (*mongo.Collection)(client.Database("URLs").Collection("col"))

	fmt.Println("Collection instance is ready")
}

func main() {
	//client, _ = mongo.Connect(ctx, "mongodb://localhost:27017")
	router := mux.NewRouter()
	router.HandleFunc("/shorten", CreateEndpoint).Methods("POST")
	router.HandleFunc("/expand/", ExpandEndpoint).Methods("POST")
	router.HandleFunc("/{id}", RedirectEndpoint).Methods("GET")
	router.HandleFunc("/", Home).Methods("GET")
	log.Fatal((http.ListenAndServe(":8000", router)))
}

type URL struct {
	LongURL  string    `json:"longURL,omitempty" bson:"_id,omitempty"`
	shortURL string    `json:"shortURL,omitempty"`
	date     time.Time `json:"date,omitempty"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("<h1>This is the homepage<h1>")))
}

func CreateEndpoint(w http.ResponseWriter, r *http.Request) {
	// cache will be good option to check if data already exists
	// we haev to check if shorten URL is existing or not
	// the problem is if there is a user generating shortURL for long One
	last_num := 0x0 // this will be the last generated url which will be retrieved from the database

	i := 1 + last_num //increment i by one
	print(i)
	//upto 18 min

}

func ExpandEndpoint(w http.ResponseWriter, r *http.Request) {

}

func RedirectEndpoint(w http.ResponseWriter, r *http.Request) {

}

//database part skipped

//skipped cluster at 14:40 to 16:00
