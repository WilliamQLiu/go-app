package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/pressly/chi"
)

type loginResource struct{}

// Routes creats a REST router for the login resource
func (rs loginResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.New)     // Prompt to create login for new users
	r.Post("/", rs.Create) // POST to create a new user

	return r
}

func (rs loginResource) New(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/login.gtpl")
	t.Execute(w, nil)
	log.Println("Log: loginResource New route")
}

func (rs loginResource) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("username:", r.Form["username"])
	fmt.Println("password:", r.Form["password"])
	log.Println("Log: loginResource Create route")
}
