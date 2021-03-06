package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Groups struct {
	ID 				primitive.ObjectID 	`json:"id" bson:"_id,omitempty"`
	GroupName 		string			   	`json:"group_name" bson:"group_name,omitempty"`
	Users			[]string				`json:"users" bson:"users"`
}

type UserIdList struct {
	UsersList []string `json:"users"`
}