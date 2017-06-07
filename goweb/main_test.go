package main_test

import (
	"log"
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
