package main

import (
	//"crypto/md5"
	"database/sql"
	//"fmt"
	//"html/template"
	//"io"
	//"log"
	//"net/http"
	//"strconv"
	//"time"
	//"github.com/pressly/chi"
	//"errors"
)

//type LoginResource struct{}

type user struct {
	ID           int    `json:"id"`
	Emailaddress string `json:"emailaddress"`
	Password     string `json:"password"`
}

func (u *user) getUser(db *sql.DB) error {
	return db.QueryRow("SELECT emailaddress, password FROM users WHERE id=$1", u.ID).Scan(&u.Emailaddress, &u.Password)
}

func (u *user) updateUser(db *sql.DB) error {
	_, err := db.Exec("UPDATE users SET emailaddress=$1, password=$2 WHERE id=$3", u.Emailaddress, u.Password, u.ID)
	return err
}

func (u *user) deleteUser(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", u.ID)
	return err
}

func (u *user) createUser(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO users(emailaddress, password) VALUES($1, $2) RETURNING id", u.Emailaddress, u.Password).Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	rows, err := db.Query(
		"SELECT id, emailaddress, password FROM users LIMIT $1 OFFSET $2", count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []user{}

	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Emailaddress, &u.Password); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

//var db *sql.DB
//
//// Routes creats a REST router for the login resource
//func (rs LoginResource) Routes() chi.Router {
//	r := chi.NewRouter()
//
//	r.Get("/", rs.New)                     // Prompt to create login for new users
//	r.Post("/", rs.Create)                 // POST to create a new user
//	r.With(paginate).Get("/list", rs.List) // GET list of existing users
//
//	return r
//}
//
//func (rs LoginResource) New(w http.ResponseWriter, r *http.Request) {
//	crutime := time.Now().Unix()
//	hash := md5.New()
//	io.WriteString(hash, strconv.FormatInt(crutime, 10))
//	token := fmt.Sprintf("%x", hash.Sum(nil))
//
//	t, _ := template.ParseFiles("templates/login.gtpl")
//	t.Execute(w, token) // pass token object to template
//	log.Println("Log: LoginResource New route")
//}
//
//func (rs LoginResource) Create(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//
//	token := r.Form.Get("token")
//	var username string = template.HTMLEscapeString(r.Form.Get("username"))
//	var password string = template.HTMLEscapeString(r.Form.Get("password"))
//
//	if token != "" {
//		// Check token validity
//		fmt.Println("Token is" + token)
//	} else {
//		// Error if no token
//		fmt.Println("No Token")
//	}
//
//	if len(username) == 0 || len(password) == 0 {
//		fmt.Println("No username or password given")
//	}
//
//	fmt.Println("username:", username)
//	fmt.Println("password:", password)
//
//}
//
//func paginate(next http.Handler) http.Handler {
//	fmt.Println("TODO: Paginate")
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		// just a stub.. some ideas are to look at URL query params for something like
//		// the page number, or the limit, and send a query cursor down the chain
//		next.ServeHTTP(w, r)
//	})
//}
//
//func (rs LoginResource) List(w http.ResponseWriter, r *http.Request) {
//	log.Println("Log: LoginResource List route")
//
//	rows, err := db.Query("SELECT * FROM users;")
//
//	if err != nil {
//		fmt.Println("Error with DB")
//		log.Fatal(err)
//	}
//
//	defer rows.Close()
//}
