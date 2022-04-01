package entities

import "time"

type User struct {
	ID string 			`json:"id,omitempty" bson:"id"`
	Name string 		`json:"name" bson:"name"`
	Username string 	`json:"username" bson:"username"`
	Email string 		`json:"email" bson:"email"`
	Phone string 		`json:"phone" bson:"phone"`
	Website string 		`json:"website" bson:"website"`
	Address Address  	`json:"address" bson:"address"`
	Company Company  	`json:"company" bson:"address"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type Address struct {
	Street string 	`json:"street" bson:"street"`
	Suite string 	`json:"suite" bson:"suite"`
	City string 	`json:"city" bson:"city"`
	Zipcode string 	`json:"zipcode" bson:"zipcode"`
	Geo Geo 	`json:"geo" bson:"geo"`
}

type Company struct {
	Name string 	`json:"name" bson:"name"`
	CatchPhrase string 	`json:"catchphrase" bson:"catchphrase"`
	Bs string 	`json:"bs" bson:"bs"`
}

type Geo struct {
	Lat string 	`json:"lat" bson:"lat"`
	Lng string 	`json:"lng" bson:"lng"`
}
