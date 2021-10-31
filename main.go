package main

import (
	"context"
	"encoding/json"
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
	log.Fatal((http.ListenAndServe(":12345", router)))
}

type URL struct {
	LongURL  string    `json:"longURL,omitempty" bson:"longURL,omitempty"`
	ShortURL string    `json:"shortURL,omitempty" bson:"shortURL,omitempty"`
	Date     time.Time `json:"date,omitempty" bson:"date,omitempty"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("<h1>This is the homepage<h1>")))
}

func CreateEndpoint(w http.ResponseWriter, r *http.Request) {
	// cache will be good option to check if data already exists
	// we have to check if shortened URL is existing or not
	// the problem is if there is a user generating shortURL for long One
	//increment i by one
	var url URL
	_ = json.NewDecoder(r.Body).Decode(&url)
	//fmt.Println(url.LongURL) // this is long url
	w.Header().Set("Content-Type", "application/json")
	url.ShortURL = "11"
	url.Date = time.Now()
	fmt.Println(url.LongURL)
	fmt.Println(url.ShortURL)
	fmt.Println(url.Date)
	insertdata(url)
	json.NewEncoder(w).Encode(url)
	//upto 18 min
}

func insertdata(new URL) {
	inserted, err := collection.InsertOne(context.Background(), new)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inserted.InsertedID)
}

func ExpandEndpoint(w http.ResponseWriter, r *http.Request) {

}

func RedirectEndpoint(w http.ResponseWriter, r *http.Request) {
	
}

//database part skipped

//skipped cluster at 14:40 to 16:00
