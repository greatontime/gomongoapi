package dao

import (
	"context"
	"fmt"
	"log"
	"github.com/greatontime/gomongoapi/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"	
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://localhost:27017"

//DBNAME Database name
const DBNAME = "phonebook"

//COLLECTIONNAME Collection name
const COLLECTIONNAME = "people"

var db *mongo.Database

//Connect establish a connection to database
func init(){
	clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
	client,err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(DBNAME)
}

//InsertManyValues inserts many items from byte slice
func InsertManyValues(people []models.Person){
	var peoplelist []interface{}
	for _, p := range people{
		peoplelist = append(peoplelist, p)
	}
	_, err := db.Collection(COLLECTIONNAME).InsertMany(context.Background(),peoplelist)
	if err != nil {
		log.Fatal(err)
	}
}

//InsertOneValue inserts one item from Person model
func InsertOneValue(person models.Person){
	fmt.Println(person)
	_, err := db.Collection(COLLECTIONNAME).InsertOne(context.Background(), person)
	if err != nil {
		log.Fatal(err)
	}
}

//GetAllPeople returns all people from DB
func GetAllPeople() []models.Person {
	cur, err := db.Collection(COLLECTIONNAME).Find(context.Background(),nil,nil)
	if err != nil {
		log.Fatal(err)
	}

	var elements []models.Person
	var elem models.Person 
	//Get the next result from the cursor

	for cur.Next(context.Background()){
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		elements = append(elements,elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return elements
}
//DeletePerson deletes an existing person
func DeletePerson(person models.Person){
	_, err := db.Collection(COLLECTIONNAME).DeleteOne(context.Background(), person, nil)
	if err != nil {
		log.Fatal(err)
	}
}

//UpdatePerson updates and existing person
func UpdatePerson(person models.Person, personID string) {
	doc := db.Collection(COLLECTIONNAME).FindOneAndUpdate(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("id", personID),
		),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				bson.EC.String("firstname", person.Firstname),
				bson.EC.String("lastname", person.Lastname),
				bson.EC.String("contactinfo.city", person.City),
				bson.EC.String("contactinfo.zipcode", person.Zipcode),
				bson.EC.String("contactinfo.phone", person.Phone)),
		),
		nil)
	fmt.Println(doc)
}
