package controller

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // _ means to import only for its side-effects (initialization)
	"github.com/williamqliu/go-app/goweb/model"
)

// App : struct to expose references to the router and database
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Parse and cache all templates
// template.Must checks for parsing errors
var templates = template.Must(template.ParseGlob("./view/*"))

// InitializeDB : connect to database
func (app *App) InitializeDB(user, password, dbname, dbhostname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	if len(dbhostname) > 0 {
		connectionString = fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", user, password, dbname, dbhostname)
	}

	var err error
	app.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

// InitializeTables : initialize database tables if not exists
func (app *App) InitializeTables() {
	fmt.Println("Initializing Table")
	if _, err := app.DB.Query(model.UserTableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

// InitializeRoutes : func to initialize routes
func (app *App) InitializeRoutes() {
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/", indexHandler).Methods("GET")
	app.Router.HandleFunc("/hello", helloHandler).Methods("GET")
	app.Router.HandleFunc("/login", app.getLoginTemplate).Methods("GET")
	app.Router.HandleFunc("/users", app.getUsers).Methods("GET")
	app.Router.HandleFunc("/users", app.createUser).Methods("POST")
	app.Router.HandleFunc("/users/{id:[0-9]+}", app.getUser).Methods("GET")
	app.Router.HandleFunc("/users/{id:[0-9]+}", app.updateUser).Methods("PUT")
	app.Router.HandleFunc("/users/{id:[0-9]+}", app.deleteUser).Methods("DELETE")

	// Serve Static Files through web server instead of reverse-proxy
	var dir string
	flag.StringVar(&dir, "dir", "./static", "the directory to serve files from, defaults to current dir") // bind string to flag with key, value, and comment
	flag.Parse()                                                                                          // call Parse() after flags are defined
	app.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
}

// Run : func to start the main application at specifc port
func (app *App) Run(addr string) {
	err := http.ListenAndServe(addr, app.Router)
	if err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// template.Must checks for parsing errors
	//tmpl, err := template.ParseFiles("./view/index.html")
	//tmpl = template.Must(template.ParseFiles("./view/content.tmpl", "./view/header.tmpl", "./view/footer.tmpl"))
	err := templates.ExecuteTemplate(w, "indexPage", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Log: indexHandler request")
}

// helloHandler : func to parse form data (e.g. url with '/hello?abcd=12')
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
