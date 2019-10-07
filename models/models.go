package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Person Struct for person model
type Person struct {
	ID        primitive.ObjectID `bson:"_id"`
	Firstname string `json:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Contactinfo `json:"contactinfo,omitempty"`
}
//Contactinfo Struct for contactinfo model
type Contactinfo struct {
	City string `json:"city,omitempty"`
	Zipcode string `json:"zipcode,omitempty"`
	Phone string `json:"phone,omitempty"`
}