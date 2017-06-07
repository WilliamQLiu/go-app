package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // _ means to import only for its side-effects (initialization)
)

// struct to expose references to the router and database
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// func to initialize database
func (a *App) InitializeDB(user, password, dbname string) {}

// func to run the application
func (a *App) Run(addr string) {}
