// Main app

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/pressly/chi"
	"github.com/williamqliu/go-app/models"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "goappdb"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
	log.Println("Log: indexHandler request")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Handler") // send data to client side

	// print the parsed form on server side
	r.ParseForm()                       // parse arguments, e.g. /hello/?url_long=111&url_long=222
	fmt.Println("path", r.URL.Path)     // `/hello/`
	fmt.Println("scheme", r.URL.Scheme) // scheme
	fmt.Println("method", r.Method)     // method GET
	fmt.Println(r.Form["url_long"])     // [111 222]
	for k, v := range r.Form {
		fmt.Println("key:", k)                    // key: url_long
		fmt.Println("val:", strings.Join(v, " ")) // val: 111 222
	}
	log.Println("Log: helloHandler request")
}

func main() {
	log.Println("Log: main app is running")

	r := chi.NewRouter()
	r.Get("/", indexHandler)
	r.Get("/hello/", helloHandler)
	r.Mount("/login", loginResource{}.Routes())

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err) // panic stops ordinary flow of control and begins panicking (program crashes)
	}
}
