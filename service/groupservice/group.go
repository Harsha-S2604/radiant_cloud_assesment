package groupservice

import (
	"net/http"
	"context"


	"radiant_cloud_assesment/models"
	"radiant_cloud_assesment/utilities/validations"
	"radiant_cloud_assesment/utilities/misc"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"

)

func AddGroupHandler(db *mongo.Database) gin.HandlerFunc {

	addGroup := func(ctx *gin.Context) {
		var group models.Groups
		ctx.ShouldBindJSON(&group)
		isValidGroupData, msg := validations.ValidateGroupData(group)
		if !isValidGroupData {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": msg,
			})
			return
		}

		isGroupExist := misc.CheckGroupExist(db, group.GroupName)
		if isGroupExist {
			ctx.JSON(http.StatusConflict, gin.H{
				"success": false,
				"message": "group already exists",
			})
			return
		}

		insertResult, insertResultErr := db.Collection("groups").InsertOne(ctx, group)
		if insertResultErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Sorry, something went wrong. our team is working on it. Please try again later",
			})
			return
		}

		groupId := insertResult.InsertedID.(primitive.ObjectID)
		ctx.JSON(http.StatusCreated, gin.H{
			"success": true,
			"id": groupId,
			"message": "group created successfully",
		})
	}

	return gin.HandlerFunc(addGroup)
}

func DeleteGroupHandler(db *mongo.Database) gin.HandlerFunc {
	
	deleteGroup := func(ctx *gin.Context) {
		groupId := ctx.Params.ByName("id")
		if groupId == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "group id is required",
			})
			return
		}
		objectId, objectIdErr := primitive.ObjectIDFromHex(groupId)
		if objectIdErr != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "invaid group id",
			})
			return
		}

		isGroupExist := misc.CheckGroupExistById(db, objectId)
		if !isGroupExist {
			ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "group not found",
			})
			return
		}

		res, deleteErr := db.Collection("groups").DeleteOne(context.TODO(), 
		bson.D{{"_id", objectId}})

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
			"message": "group deleted successfully",
		})
	}
	return gin.HandlerFunc(deleteGroup)
}