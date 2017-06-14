package controller_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/williamqliu/go-app/goweb/controller"
	"github.com/williamqliu/go-app/goweb/model"
	"github.com/williamqliu/go-app/goweb/util"
)

var app controller.App

func TestMain(m *testing.M) {
	app = controller.App{}
	app.InitializeDB(
		"postgres", // dbUsername
		"postgres", // dbPassword
		"postgres", // dbName
		"")         // dbHostName
	app.InitializeRoutes()
	ensureTableExists()
	code := m.Run() // Tests are executed with this
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	log.Println("Ensuring that Table exists")
	if _, err := app.DB.Exec(model.UserTableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	app.DB.Exec("DELETE FROM users")
	app.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	// Executes the request using the application's router and returns response
	responseRecorder := httptest.NewRecorder()
	app.Router.ServeHTTP(responseRecorder, req)
	return responseRecorder
}

func addUser(count int) {
	// Add one or more records into the table for testing
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		app.DB.Exec("INSERT INTO users(emailaddress, password) VALUES($1, $2)", "testuser@williamqliu.com", "test12345")
	}
}

func TestLoginTemplate(t *testing.T) {
	// Check that login routing works
	clearTable()

	os.Chdir("..") // test creates a _test dir, need to move up one dir to find template
	req, _ := http.NewRequest("GET", "/login", nil)
	response := executeRequest(req)

	util.CheckResponseCode(t, http.StatusOK, response.Code)
}

func TestEmptyTable(t *testing.T) {
	// Deletes all records from 'users' table, send GET request to /users endpoint
	clearTable()

	req, _ := http.NewRequest("GET", "/users", nil)
	response := executeRequest(req)

	util.CheckResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentUser(t *testing.T) {
	// Get an nonexistent user, check for 404 Not Found and contains error message
	clearTable()

	req, _ := http.NewRequest("GET", "/users/999999", nil)
	response := executeRequest(req)

	util.CheckResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "User not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", m["error"])
	}
}

func TestCreateUser(t *testing.T) {
	// Create a user, returns 201 and right key-values in response body payload
	clearTable()

	payload := []byte(`{"emailaddress":"testuser@williamqliu.com", "password": "test12345"}`)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	response := executeRequest(req)

	util.CheckResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["emailaddress"] != "testuser@williamqliu.com" {
		t.Errorf("Expected user name to be 'testuser@williamqliu.com'. Got '%v'", m["emailaddress"])
	}

	if m["password"] != "test12345" {
		t.Errorf("Expected user password to be 'test12345'. Got '%v'", m["password"])
	}
}

func TestGetUser(t *testing.T) {
	// Add a User and tests retrieving user returns 200 OK
	clearTable()

	addUser(1)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)

	util.CheckResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateUser(t *testing.T) {
	// Update an existing User data
	clearTable()
	addUser(1)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)

	var originalUser map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalUser)

	payload := []byte(`{"emailaddress":"newtestuser@williamqliu.com", "password": "newtest12345"}`)

	req, _ = http.NewRequest("PUT", "/users/1", bytes.NewBuffer(payload))
	response = executeRequest(req)

	util.CheckResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalUser["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalUser["id"], m["id"])
	}

	if m["emailaddress"] == originalUser["emailaddress"] {
		t.Errorf("Expected the emailaddress to change from '%v' to '%v'. Got '%v'", originalUser["emailaddress"], m["emailaddress"], m["emailaddress"])
	}

	if m["password"] == originalUser["password"] {
		t.Errorf("Expected the password to change from '%v' to '%v'. Got '%v'", originalUser["password"], m["password"], m["password"])
	}
}

func TestDeleteUser(t *testing.T) {
	clearTable()
	addUser(1)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)
	util.CheckResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/users/1", nil)
	response = executeRequest(req)

	util.CheckResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/users/1", nil)
	response = executeRequest(req)
	util.CheckResponseCode(t, http.StatusNotFound, response.Code)
}
