package main

import (
	"os"

	"radiant_cloud_assesment/routes"
	"radiant_cloud_assesment/config/db"
)

func main() {
	/*
		1. database connection(MongoDB)
		2. setup router(Gin)
		3. start server
	*/
	database := db.ConnectDB()
	r := routes.SetupRouter(database)
	port := os.Getenv("PORT")
	r.Run(":"+port)
}