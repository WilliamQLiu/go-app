package main_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/williamqliu/go-app/goweb"
)

var app main.App

func TestMain(m *testing.M) {
	app = main.App{}
	app.InitializeDB(
		"postgres", // DB_Username
		"postgres", // DB_Password
		"postgres") // DB_Name

	ensureTableExists()

	code := m.Run() // Tests are executed with this

	clearTable()

	os.Exit(code)
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS users
(
	id INTEGER PRIMARY KEY,
	emailaddress varchar(254) NOT NULL,
	password varchar(254) NOT NULL
)`

func ensureTableExists() {
	if _, err := app.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	app.DB.Exec("DELETE FROM users")
	//app.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	// Executes the request using the application's router and returns response
	responseRecorder := httptest.NewRecorder()
	app.Router.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

func checkResponseCode(t *testing.T, expected, actual int) {
	// Compares response code
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestEmptyTable(t *testing.T) {
	// Deletes all records from 'users' table, send GET request to /users endpoint
	clearTable()

	req, _ := http.NewRequest("GET", "/users", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentUser(t *testing.T) {
	// Get an nonexistent user, check for 404 Not Found and contains error message
	clearTable()

	req, _ := http.NewRequest("GET", "/user/999999", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "User not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", m["error"])
	}
}

func TestCreateUser(t *testing.T) {
	// Create a user, returns 201
	clearTable()

	payload := []byte(`{"emailaddress":"testuser@williamqliu.com", "password": "test12345"}`)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["emailaddress"] != "testuser@williamqliu.com" {
		t.Errorf("Expected user name to be 'testuser@williamqliu.com'. Got '%v'", m["emailaddress"])
	}

	if m["password"] != "test12345" {
		t.Errorf("Expected user password to be 'test12345'. Got '%v'", m["password"])
	}
}
