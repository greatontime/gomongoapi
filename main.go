package main
import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/greatontime/gomongoapi/handlers"
	"github.com/greatontime/gomongoapi/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBNAME Database name
const DBNAME = "phonebook"

//COLLECTION Collection name
const COLLECTION = "people"

//CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://localhost:27017"

func init(){
	//Populates database with dummy data

	var people []models.Person 
	clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
	client,err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(DBNAME)

	//Load values from json file to model
	byteValues, err := ioutil.ReadFile("person_data.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(byteValues, &people)

	//Insert people to Database
	var peoplelist []interface{}
	for _, p := range people{
		peoplelist = append(peoplelist,p)
	}
	_, err = db.Collection(COLLECTION).InsertMany(context.Background(),peoplelist)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", handlers.GetAllPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", handlers.GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people", handlers.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people", handlers.DeletePersonEndpoint).Methods("DELETE")
	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
