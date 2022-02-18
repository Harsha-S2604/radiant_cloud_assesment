package userservice

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

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	jsonData, err := json.Marshal(map[string]interface{}{
		"email": "test@gmail.com",
		"first_name": "test",
		"last_name": "user",
		"userid": "tuser",
	})

	if err != nil {
		t.Fail()
	}

	database := db.ConnectDB()
	req, err := http.NewRequest(http.MethodPost, "api/v1/users", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fail()
	}

	req.Header.Set("Content-type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	addUser := AddUserHandler(database)
	addUser(c)
	bodySb, err := ioutil.ReadAll(w.Body)
	var decodedResponse interface{}
	err = json.Unmarshal(bodySb, &decodedResponse)
	if err != nil {
		t.Fatalf("Cannot decode response <%p> from server. Err: %v", bodySb, err)
	}

	expected, actual := 201, w.Code
	if expected == actual {
		t.Logf("Create User Success")
	} else {
		t.Fail()
	}
}

func TestGetUserById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database := db.ConnectDB()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{
			Key: "id",
			Value: "tuser", // change the user id you want
		},
	}

	getUserById := GetUserByIdHandler(database)
	getUserById(c)

	expected, actual := 302, w.Code // change the code according to the user id
	if expected != actual {
		t.Errorf("Expected code %d but got %d", expected, actual)
	}

}

func TestUpdateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	jsonData, err := json.Marshal(map[string]interface{}{
		"email": "test123@gmail.com",
	})

	if err != nil {
		t.Fail()
	}
	database := db.ConnectDB()
	req, err := http.NewRequest(http.MethodPost, "api/v1/users", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fail()
	}

	req.Header.Set("Content-type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = []gin.Param{
		{
			Key: "id",
			Value: "tuser", // change the user id you want
		},
	}

	updateUser := UpdateUserHandler(database)
	updateUser(c)
	expected, actual := 200, w.Code
	if expected != actual {
		t.Errorf("Expected code %d but got %d", expected, actual)
	}

}

func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database := db.ConnectDB()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{
			Key: "id",
			Value: "tuser", // change the user id you want
		},
	}

	deleteUser := DeleteUserHandler(database)
	deleteUser(c)
	expected, actual := 200, w.Code // change the code according to the user id
	if expected != actual {
		t.Errorf("Expected code %d but got %d", expected, actual)
	}
}