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

func CheckGroupExist(db *mongo.Database, groupName string) (bool) {
	var result primitive.M 
	resultErr:= db.Collection("groups").FindOne(context.TODO(), bson.M{"group_name": groupName}).Decode(&result)

	if resultErr != nil {
		fmt.Println("result error", resultErr.Error())
		return false
	}

	if result["group_name"] == "" {
		fmt.Println("result error")
		return false
	}

	

	return true
}

func CheckGroupExistById(db *mongo.Database, groupId primitive.ObjectID) (bool) {
	var result primitive.M 
	resultErr:= db.Collection("groups").FindOne(context.TODO(), bson.M{"_id": groupId}).Decode(&result)

	if resultErr != nil {
		fmt.Println("result error", resultErr.Error())
		return false
	}

	if result["group_name"] == "" {
		return false
	}

	return true
}