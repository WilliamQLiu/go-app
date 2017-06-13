package controller_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/williamqliu/go-app/goweb/controller"
	"github.com/williamqliu/go-app/goweb/model"
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

func checkResponseCode(t *testing.T, expected, actual int) {
	// Compares response code
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestLoginTemplate(t *testing.T) {
	// Check that login routing works
	clearTable()

	os.Chdir("..") // test creates a _test dir, need to move up one dir to find template
	req, _ := http.NewRequest("GET", "/login", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}
