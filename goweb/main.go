// Main app

package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"

	"app/controller"
	"app/util"
)

func main() {
	log.Println("Log: main starting to initialize app") // Log statements appear on Docker

	dbUsername := util.GetKey("dbUsername", "postgres")
	dbPassword := util.GetKey("dbPassword", "postgres")
	dbName := util.GetKey("dbPassword", "postgres")
	dbhostName := util.GetKey("hostName", "postgres") // alias created for db from docker-compose

	app := App{}
	app.Initialize(dbUsername, dbPassword, dbName, dbhostName)
	log.Println("Log: App initialized")
	app.Run(":8080")
}
