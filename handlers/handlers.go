package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/greatontime/gomongoapi/dao"
	"github.com/greatontime/gomongoapi/models"

	"github.com/gorilla/mux"
)

var people []models.Person

//GetPersonEndpoint gets a person
func GetPersonEndpoint(w http.ResponseWriter,r *http.Request){
	params := mux.Vars(r)
	payload := dao.GetAllPeople()
	for _, p := range payload{
		if p.ID == params["id"]{
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode("Person not found")
}
//GetAllPeopleEndpoint gets all people
func GetAllPeopleEndpoint(w http.ResponseWriter, r *http.Request){
	payload := dao.GetAllPeople()
	json.NewEncoder(w).Encode(payload)
}
//CreatePersonEndpoint create a person
func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request){
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	dao.InsertOneValue(person)
	json.NewEncoder(w).Encode(person)
}
//DeletePersonEndpoint delete a person
func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request){
	var person models.Person 
	_ = json.NewDecoder(r.Body).Decode(&person)
	dao.DeletePerson(person)
}
//UpdatePersonEndpoint update a person
func UpdatePersonEndpoint(w http.ResponseWriter, r *http.Request){
	personID := mux.Vars(r)["id"]
	var person models.Person 
	_ = json.NewDecoder(r.Body).Decode(&person)
	dao.UpdatePerson(person,personID)
}