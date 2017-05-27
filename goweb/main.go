// Main app

package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I'm Will.")
	log.Println("Log: handler request")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello again")
	log.Println("Log: handler request")
}

func main() {
	log.Println("Log: main app is running")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello/", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
