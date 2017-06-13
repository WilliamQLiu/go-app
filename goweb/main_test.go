package main_test

import (
	//"bytes"
	//"encoding/json"
	//"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/williamqliu/go-app/goweb/controller"
	"github.com/williamqliu/go-app/goweb/util"
)

var app controller.App

// TestMain is auto created when running `go test` unless explicitly created like below, provides a global hook to perform setup and shutdown
func TestMain(m *testing.M) {
	app = controller.App{}
	app.InitializeDB(
		"postgres", // dbUsername
		"postgres", // dbPassword
		"postgres", // dbName
		"")         // dbHostName
	app.InitializeRoutes()
	code := m.Run() // Tests are executed with this
	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	// Executes the request using the application's router and returns response
	responseRecorder := httptest.NewRecorder()
	app.Router.ServeHTTP(responseRecorder, req)
	return responseRecorder
}

func TestStaticFileServes(t *testing.T) {
	req, _ := http.NewRequest("GET", "/static/", nil)
	response := executeRequest(req)
	util.CheckResponseCode(t, http.StatusOK, response.Code)
}

func TestIndexHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req)
	util.CheckResponseCode(t, http.StatusOK, response.Code)
}
