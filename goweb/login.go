package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

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
	crutime := time.Now().Unix()
	hash := md5.New()
	io.WriteString(hash, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", hash.Sum(nil))

	t, _ := template.ParseFiles("templates/login.gtpl")
	t.Execute(w, token) // pass token object to template
	log.Println("Log: loginResource New route")
}

func (rs loginResource) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	token := r.Form.Get("token")
	var username string = template.HTMLEscapeString(r.Form.Get("username"))
	var password string = template.HTMLEscapeString(r.Form.Get("password"))

	if token != "" {
		// Check token validity
		fmt.Println("Token is" + token)
	} else {
		// Error if no token
		fmt.Println("No Token")
	}

	if len(username) == 0 || len(password) == 0 {
		fmt.Println("No username or password given")
	}

	fmt.Println("username:", username)
	fmt.Println("password:", password)
	log.Println("Log: loginResource Create route")

}
