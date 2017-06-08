// Main app

package main

import (
	//"fmt"
	//"html/template"
	"log"
	//"net/http"

	_ "github.com/lib/pq"
	//"os"
	//"strings"
	//"github.com/pressly/chi"
)

const (
	dbUsername = "postgres" // lowerCamelCase for private const variables
	dbPassword = "postgres"
	dbName     = "postgres"
)

func main() {
	log.Println("Log: main app is running")

	app := App{}
	app.Initialize(dbUsername, dbPassword, dbName)
	app.Run(":8080")

	// Setup DB
	//dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	//	DB_USER, dbPassword, dbName)
	//db, dberr := sql.Open("postgres", dbinfo)
	//checkErr(dberr)

	//defer db.Close() // defer execution of closing DB until after surround function (main) closes

	//r := chi.NewRouter()
	//r.Get("/", indexHandler)
	//r.Get("/hello/", helloHandler)
	//r.Mount("/login", LoginResource{}.Routes())

	//err := http.ListenAndServe(":8080", r)
	//if err != nil {
	//	log.Fatal(err)
	//}
}

func checkErr(err error) {
	if err != nil {
		panic(err) // panic stops ordinary flow of control and begins panicking (program crashes)
	}
}
