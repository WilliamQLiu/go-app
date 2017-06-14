package controller

import (
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/williamqliu/go-app/goweb/model"
	"github.com/williamqliu/go-app/goweb/util"
)

// getLoginTemplate : display template for login existing user
func (app *App) getLoginTemplate(w http.ResponseWriter, r *http.Request) {
	start := 0
	count := 5

	users, err := model.GetUsers(app.DB, start, count)

	err = templates.ExecuteTemplate(w, "loginPage", users)
	if err != nil {
		log.Fatal(err)
	}
}

// getSignUpTemplate : display template for signing up new user
func (app *App) getSignupTemplate(w http.ResponseWriter, r *http.Request) {
	// Create Token each visit
	crutime := time.Now().Unix()
	hash := md5.New()
	io.WriteString(hash, strconv.FormatInt(crutime, 10)) // base 10
	token := fmt.Sprintf("%x", hash.Sum(nil))

	err := templates.ExecuteTemplate(w, "signupPage", token)
	if err != nil {
		log.Fatal(err)
	}
}

// getUsers : get list of users with JSON response
func (app *App) getUsers(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	users, err := model.GetUsers(app.DB, start, count)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondWithJSON(w, http.StatusOK, users)
}

// getUser : get a single existing user with JSON response
func (app *App) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	u := model.User{ID: id}
	if err := u.GetUser(app.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			util.RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	util.RespondWithJSON(w, http.StatusOK, u)
}

// createUser : create a new user with JSON response
func (app *App) createUser(w http.ResponseWriter, r *http.Request) {
	var u model.User

	// Parse Data
	r.ParseForm()
	token := r.Form.Get("token")
	var username = template.HTMLEscapeString(r.Form.Get("username"))
	var password = template.HTMLEscapeString(r.Form.Get("password"))

	if len(username) == 0 || len(password) == 0 {
		log.Println("No username or password given")
	}

	// Check token validity
	if token != "" {
		log.Println("Token is: " + token)
	} else {
		log.Println("No Token")
	}

	log.Println("username is: ", username)
	log.Println("password is: ", password)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := u.CreateUser(app.DB); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, u)
}

// updateUser : update an existing user with JSON response
func (app *App) updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	var u model.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	u.ID = id

	if err := u.UpdateUser(app.DB); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, u)
}

// deleteUser : deletes a user with JSON response
func (app *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	u := model.User{ID: id}
	if err := u.DeleteUser(app.DB); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
