// Main app

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/pressly/chi"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I'm Will.")
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

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/login.gtpl")
		t.Execute(w, nil)
		log.Println("Log: loginHandler GET request")
	} else if r.Method == "POST" {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		log.Println("Log: loginHandler POST request")
	} else {
		fmt.Fprintf(w, "Method type not supposed")
	}
}

func main() {
	log.Println("Log: main app is running")

	r := chi.NewRouter()
	r.Get("/", indexHandler)
	r.Get("/hello/", helloHandler)
	r.Get("/login/", loginHandler)
	r.Post("/login/", loginHandler)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
