package models

type Users struct {
	Userid 			string 				`json:"userid" bson:"userid,omitempty"`
	FirstName 		string			   	`json:"first_name" bson:"first_name,omitempty"`
	LastName		string				`json:"last_name" bson:"last_name,omitempty"`
	Email			string				`json:"email" bson"email,omitempty"`
}

type UpdateUser struct {
	FirstName	string `json:"first_name"`
	LastName	string `json:"last_name"`
	Email		string `json:"email"`
}