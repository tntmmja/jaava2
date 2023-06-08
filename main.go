package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/tntmmja/jaava2/backend/config"
	"github.com/tntmmja/jaava2/backend/data"

	"github.com/tntmmja/jaava2/backend/handlers"
	//"github.com/tntmmja/jaava2/backend/utils"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("indexhandler", r.RequestURI)
	//fmt.Fprintf(w, "testing backend to frontend")
	tmpl, err := template.ParseFiles("./clientfrontend/templates/index.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Println("indexhandle2")
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("websoket")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// handle WebSocket connection here
}
func showRegistrationForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./clientfrontend/templates/register.html")
}

func showLoginForm(w http.ResponseWriter, r *http.Request) {
    // Load the login form template
    tpl, err := template.ParseFiles("./clientfrontend/templates/login.html")
    if err != nil {
        log.Println(err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Render the login form template
    err = tpl.Execute(w, nil)
    if err != nil {
        log.Println(err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}


func main() {
	r := mux.NewRouter()
	SetRoutes(r)

	http.Handle("/", r)
	//config.DBConn()
	//testing database connection, delete afterwards
	db, err := config.DBConn()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Attempt to ping the database
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	fmt.Println("Connected to the database!")
	//

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8083"},
	})
	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe("localhost:8082", handler))
}

var SetRoutes = func(router *mux.Router) {
	router.HandleFunc("/socket", handleWebSocket)
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/register", data.RegisterHandler).Methods(("POST"))
	router.HandleFunc("/register", showRegistrationForm).Methods("GET")
	router.HandleFunc("/login", handlers.LoginHandler).Methods(("POST"))
	router.HandleFunc("/login", showLoginForm).Methods("GET")
	//router.HandleFunc("/loggedUser", handlers.LoggedHandler)
}
