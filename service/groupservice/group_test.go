package groupservice

import (
	"testing"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"bytes"

	"radiant_cloud_assesment/config/db"

	"github.com/gin-gonic/gin"
)

func TestCreateGroup(t *testing.T) {
	gin.SetMode(gin.TestMode)
	jsonData, err := json.Marshal(map[string]interface{}{
		"group_name": "developer",
		"users": []string{"rmike", "djackson"},
	})

	if err != nil {
		t.Fail()
	}

	database := db.ConnectDB()
	req, err := http.NewRequest(http.MethodPost, "api/v1/groups", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fail()
	}

	req.Header.Set("Content-type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	addGroup := AddGroupHandler(database)
	addGroup(c)
	bodySb, err := ioutil.ReadAll(w.Body)
	var decodedResponse interface{}
	err = json.Unmarshal(bodySb, &decodedResponse)
	if err != nil {
		t.Fatalf("Cannot decode response <%p> from server. Err: %v", bodySb, err)
	}

	expected, actual := 201, w.Code
	if expected == actual {
		t.Logf("Create group success")
	} else {
		t.Fail()
	}
}

func TestUpdateGroup(t *testing.T) {
	gin.SetMode(gin.TestMode)
	jsonData, err := json.Marshal(map[string]interface{}{
		"users": []string{"troy", "jgyllenhall", "ealderson"},
	})

	if err != nil {
		t.Fail()
	}
	database := db.ConnectDB()
	req, err := http.NewRequest(http.MethodPut, "api/v1/groups/:group_name", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fail()
	}

	req.Header.Set("Content-type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = []gin.Param{
		{
			Key: "group_name",
			Value: "developer", // change the group name you want
		},
	}

	updateGroup := UpdateGroupHandler(database)
	updateGroup(c)
	expected, actual := 200, w.Code
	if expected != actual {
		t.Errorf("Expected code %d but got %d", expected, actual)
	}

}

func TestGetGroupUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database := db.ConnectDB()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{
			Key: "group_name",
			Value: "developer", // change the group name you want
		},
	}

	getGroupUsers := GetGroupUsersHandler(database)
	getGroupUsers(c)

	expected, actual := 302, w.Code // change the code according to the group name
	if expected != actual {
		t.Errorf("Expected code %d but got %d", expected, actual)
	}
}

func TestDeleteGroup(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database := db.ConnectDB()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{
			Key: "group_name",
			Value: "developer", // change the group id you want
		},
	}

	deleteGroup := DeleteGroupHandler(database)
	deleteGroup(c)
	expected, actual := 200, w.Code // change the code according to the group id
	if expected != actual {
		t.Errorf("Expected code %d but got %d", expected, actual)
	}
}