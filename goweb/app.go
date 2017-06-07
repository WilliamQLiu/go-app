package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // _ means to import only for its side-effects (initialization)
	"log"
)

// struct to expose references to the router and database
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// func to initialize database
func (app *App) InitializeDB(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	app.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter()
}

// func to run the application
func (a *App) Run(addr string) {}
