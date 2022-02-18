package userservice

import (
	"net/http"
	"log"
	"context"

	"radiant_cloud_assesment/models"
	"radiant_cloud_assesment/utilities/validations"
	"radiant_cloud_assesment/utilities/misc"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUserHandler(db *mongo.Database) gin.HandlerFunc {

	addUser := func(ctx *gin.Context) {
		/*
			1. bind request body with user object of type models.Users
			2. validate user email by calling ValidateUserEmail from package validations
			3. check if user exists
				if exist 
					don't add user to db and exit the program
				else
			 		add user to database

			data structure for users
			{
				"userid": "tuser",
				"first_name": "test",
				"last_name": "user",
				"email": "test@gmail.com",
				"groups": []
			}
		*/
		var user models.Users
		ctx.ShouldBindJSON(&user)
		isValidUserData, msg := validations.ValidateUserData(user)
		if !isValidUserData {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": msg,
			})
			return
		}
		isUserExist, msg := misc.CheckUserExist(db, user)
		if isUserExist {
			ctx.JSON(http.StatusConflict, gin.H{
				"success": false,
				"message": msg,
			})
			return
		}

		_, insertResultErr := db.Collection("users").InsertOne(ctx, user)
		if insertResultErr != nil {
			log.Println("add user error:", insertResultErr.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Sorry, something went wrong. our team is working on it. Please try again later",
			})
			return
		}
		
		ctx.JSON(http.StatusCreated, gin.H{
			"success": true,
			"id": user.Userid,
			"message": "User registration successful. Please sign-in to continue.",
		})
	}

	return gin.HandlerFunc(addUser)

}

func GetUserByIdHandler(db *mongo.Database) gin.HandlerFunc {

	getUserById := func (ctx *gin.Context) {
		/*
			1. get user_id from the url params
			2. get the user by querying the db
				if exist
					send success message along with user data
				else
					send user not found message  
		*/
		userId := ctx.Params.ByName("id")
		if userId == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "user id is required",
			})
			return
		}
		
		var user primitive.M
		_ = db.Collection("users").FindOne(context.TODO(), bson.M{"userid": userId}).Decode(&user)
		if userid, ok := user["userid"]; !ok || userid == "" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "user not found",
			})
			return
		}

		ctx.JSON(http.StatusFound, gin.H{
			"success": true,
			"message": "user found",
			"userData": user,
		})
	}
	return gin.HandlerFunc(getUserById)
}

func UpdateUserHandler(db *mongo.Database) gin.HandlerFunc {

	updateUser := func(ctx *gin.Context) {
		userId := ctx.Params.ByName("id")
		if userId == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "user id is required",
			})
			return
		}

		var userFromDB models.Users
		_ = db.Collection("users").FindOne(context.TODO(), bson.M{"userid": userId}).Decode(&userFromDB)
		if userFromDB.Email == "" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "user not found",
			})
			return
		}

		var updateUserObj models.UpdateUser
		ctx.ShouldBindJSON(&updateUserObj)
		if updateUserObj.FirstName == "" {
			updateUserObj.FirstName = userFromDB.FirstName
		}

		if updateUserObj.LastName == "" {
			updateUserObj.LastName = userFromDB.LastName
		}

		if updateUserObj.Email == "" {
			updateUserObj.Email = userFromDB.Email
		}

		var groupId primitive.ObjectID
		var isGroupExist bool
		var groups []models.Groups
		if updateUserObj.GroupName != "" {
			isGroupExist, groupId = misc.CheckGroupExist(db, updateUserObj.GroupName)
			if isGroupExist {
				isUserInTheGroup := misc.CheckUserInTheGroup(db, updateUserObj.GroupName, userId)
				if !isUserInTheGroup {
					groups = userFromDB.Groups
					groups = append(groups, models.Groups{
						groupId,
						updateUserObj.GroupName,
					})
				}
			}
		}

		if len(groups) == 0 {
			groups = userFromDB.Groups
		}
		updateResult, updateErr := db.Collection("users").UpdateOne(
			context.TODO(), bson.M{"userid": userId},
			bson.D{
				{"$set", bson.D{{"groups", groups},
				{"first_name", updateUserObj.FirstName},
				{"last_name", updateUserObj.LastName},
				{"email", updateUserObj.Email}}},
			},
		)

		if updateErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Sorry, something went wrong. Please try again later.",
			})
			return
		}

		if !(updateResult.ModifiedCount > 0) {
			ctx.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "No updates were done",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Updated successfully.",
		})
	}

	return gin.HandlerFunc(updateUser)
}

func DeleteUserHandler(db *mongo.Database) gin.HandlerFunc {
	
	deleteUser := func(ctx *gin.Context) {
		userId := ctx.Params.ByName("id")
		if userId == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "user id is required",
			})
			return
		}
		isUserExist := misc.CheckUserExistById(db, userId)
		if !isUserExist {
			ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "user not found",
			})
			return
		}

		res, deleteErr := db.Collection("users").DeleteOne(context.TODO(), 
		bson.D{{"userid", userId}})

		if deleteErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Sorry, something went wrong. Please try again later",
			})
			return
		}

		if !(res.DeletedCount > 0) {
			ctx.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "No deletions were made",
			})
			return
		}


		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "user deleted successfully",
		})
	}
	return gin.HandlerFunc(deleteUser)
}