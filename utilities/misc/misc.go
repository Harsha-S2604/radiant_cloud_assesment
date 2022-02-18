package misc

import (
	"context"
	"fmt"

	"radiant_cloud_assesment/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckUserExist(db *mongo.Database, user models.Users) (bool, string) {
	var result primitive.M 
	resultErr:= db.Collection("users").FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&result)
	if resultErr != nil {
		return false, ""
	}

	if result["email"] == "" {
		return false, ""
	}

	return true, "user already exist"
}

func CheckUserExistById(db *mongo.Database, userId string) (bool) {
	var result primitive.M 
	resultErr:= db.Collection("users").FindOne(context.TODO(), bson.M{"userid": userId}).Decode(&result)
	if resultErr != nil {
		return false
	}

	if result["email"] == "" {
		return false
	}

	return true
}

func CheckGroupExist(db *mongo.Database, groupName string) (bool, primitive.ObjectID) {
	var result primitive.M 
	resultErr:= db.Collection("groups").FindOne(context.TODO(), bson.M{"group_name": groupName}).Decode(&result)

	idStr := fmt.Sprintf("%v", result["_id"])
	objectId, objectIdErr := primitive.ObjectIDFromHex(idStr)
	if objectIdErr != nil{
		return false,  objectId
	}

	if resultErr != nil {
		return false, objectId
	}

	if result["group_name"] == "" {
		return false, objectId
	}

	

	return true, objectId
}

func CheckUserInTheGroup(db *mongo.Database, groupName string, userId string) (bool) {
	var result models.Users 
	resultErr:= db.Collection("users").FindOne(context.TODO(), bson.M{"userid": userId}).Decode(&result)
	if resultErr != nil {
		return false
	}

	groups := result.Groups

	for _, val := range groups {
		fmt.Println(val)
	}

	for i := 0; i < len(groups); i++ {
		if groups[i].GroupName == groupName {
			return true
		}  
	}

	return false
}