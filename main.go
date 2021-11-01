package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var coll *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//ctx, _ := context.WithTimeout(context.Background(), 10+time.Second)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	collection = (*mongo.Collection)(client.Database("URLs").Collection("col"))
	coll = (*mongo.Collection)(client.Database("Last").Collection("row"))
	fmt.Println("Collection instance is ready")
}

func main() {
	//client, _ = mongo.Connect(ctx, "mongodb://localhost:27017")
	router := mux.NewRouter()
	router.HandleFunc("/shorten", CreateEndpoint).Methods("POST")
	router.HandleFunc("/{id}", RedirectEndpoint).Methods("GET")
	router.HandleFunc("/", Home).Methods("GET")
	log.Fatal((http.ListenAndServe(":12345", router)))
}

type URL struct {
	LongURL  string    `json:"longURL,omitempty" bson:"longURL,omitempty"`
	ShortURL string    `json:"shortURL,omitempty" bson:"shortURL,omitempty"`
	Date     time.Time `json:"date,omitempty" bson:"date,omitempty"`
}

type LastURL struct {
	Last int64 `json:"lastURL,omitempty" bson:"lastURL, omitempty"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("<h1>This is the homepage<h1>")))
}

func CreateEndpoint(w http.ResponseWriter, r *http.Request) {
	// cache will be good option to check if data already exists
	// we have to check if shortened URL is existing or not
	// the problem is if there is a user generating shortURL for long One
	//increment i by one
	var last LastURL
	err := coll.FindOne(context.TODO(), bson.M{"lastURL": 0}).Decode(&last)
	if err != nil {
		fmt.Print(err)
		//fmt.Println("line65")
		var new LastURL
		new.Last = 0
		inserted, err := coll.InsertOne(context.Background(), new)
		fmt.Println(inserted)
		if err != nil {
			print(err)
		}
		return
	}
	// if last is found then do update else make a new one and send it ot frontends
	fmt.Println(last)
	var url URL
	_ = json.NewDecoder(r.Body).Decode(&url)
	//fmt.Println(url.LongURL) // this is long url
	w.Header().Set("Content-Type", "application/json")
	url.ShortURL = "1111"
	url.Date = time.Now()
	fmt.Println(url.LongURL)
	fmt.Println(url.ShortURL)
	fmt.Println(url.Date)
	insertdata(url)
	json.NewEncoder(w).Encode(url)
}

func insertdata(new URL) {
	inserted, err := collection.InsertOne(context.Background(), new)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inserted.InsertedID)
}

func findurl(short string) string {
	var long URL
	fmt.Println("line 80" + short)
	get, _ := primitive.ObjectIDFromHex(short)
	fmt.Print(get)
	fmt.Println("line84")
	err := collection.FindOne(context.TODO(), bson.M{"shortURL": short}).Decode(&long)

	if err != nil {
		fmt.Println(err)
	}
	return long.LongURL
}

func ExpandEndpoint(w http.ResponseWriter, r *http.Request) {

}

//this is done
func RedirectEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["id"] + "line98")
	fmt.Println("this reached line 99")
	long := findurl(params["id"])
	//json.NewEncoder(w).Encode(long)
	http.Redirect(w, r, long, http.StatusPermanentRedirect)
}
