package routes

import (
	"testing"

	"radiant_cloud_assesment/config/db"

	"github.com/gin-gonic/gin"
)

func TestInitializeRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database := db.ConnectDB()
	routeHandler := SetupRouter(database)
	if routeHandler == nil {
		t.Fail()
	}
}