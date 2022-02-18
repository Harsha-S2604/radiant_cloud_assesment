package routes

import (
	"radiant_cloud_assesment/service/userservice"
	"radiant_cloud_assesment/service/groupservice"

	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func SetupRouter(db *mongo.Database) *gin.Engine {
	
	router := gin.Default()
	router.Use(cors.Default())

	userAPIRouter := router.Group("api/v1")
	{

		userAPIRouter.GET("/users/:id", userservice.GetUserByIdHandler(db))

		userAPIRouter.POST("/users", userservice.AddUserHandler(db))

		userAPIRouter.PUT("/users/:id", userservice.UpdateUserHandler(db))

		userAPIRouter.DELETE("/users/:id", userservice.DeleteUserHandler(db))

	}

	groupAPIRouter := router.Group("api/v1")
	{
		groupAPIRouter.POST("/groups", groupservice.AddGroupHandler(db))

	}

	return router

}