// Main app

package main

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/williamqliu/go-app/goweb/controller"
	"github.com/williamqliu/go-app/goweb/util"
)

// main loads configuration settings, registers database, and runs the server
func main() {
	log.Println("Log: main starting to initialize app") // Log statements appear on Docker

	dbUsername := util.GetKey("dbUsername", "postgres")
	dbPassword := util.GetKey("dbPassword", "postgres")
	dbName := util.GetKey("dbPassword", "postgres")
	dbhostName := util.GetKey("hostName", "postgres") // alias created for db from docker-compose

	app := controller.App{}
	app.InitializeDB(dbUsername, dbPassword, dbName, dbhostName)
	app.InitializeTables()
	app.InitializeRoutes()
	log.Println("Log: App initialized!")
	app.Run(":8080")
}
