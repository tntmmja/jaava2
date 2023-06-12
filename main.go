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
	//
)

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


	// Create a new WebSocketManager instance
	manager := config.NewWebSocketManager()

	// Start the WebSocket manager
	go manager.Run()

	http.Handle("/", r)
	//config.DBConn()
	//testing database connection, delete afterwards??
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

	fmt.Println("Connected main to the database!")

	// Initialize and start the WebSocket manager
	WebSocketManager := config.NewWebSocketManager()
	go WebSocketManager.Run()


	//probably cors is not needed if vue is not used

	// c := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"http://localhost:8083"},
	// })
	// handler := c.Handler(r)

	//log.Fatal(http.ListenAndServe("localhost:8082", handler))
	// Start the HTTP server
	log.Println("Server is running on http://localhost:8082")
	err2 := http.ListenAndServe(":8082", nil)
	if err2 != nil {
		log.Fatal("Server failed:", err2)
	}

}

var SetRoutes = func(router *mux.Router) {
	router.HandleFunc("/socket", config.HandleWebSocket)
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/register", data.RegisterHandler).Methods(("POST"))
	router.HandleFunc("/register", showRegistrationForm).Methods("GET")
	router.HandleFunc("/login", handlers.LoginHandler).Methods(("POST"))
	router.HandleFunc("/login", showLoginForm).Methods("GET")
	router.HandleFunc("/create-post", handlers.CreatePostHandler).Methods("POST")
	// router.HandleFunc("/create-comment", handlers.CreateCommentHandler).Methods("POST")
	router.HandleFunc("/feed", handlers.FeedHandler).Methods("GET", "POST")
	// router.HandleFunc("/filter", handlers.FilterHandler).Methods("GET") // Add this line for the FilterHandler
	router.HandleFunc("/post/{id}", handlers.PostHandler).Methods("GET")
	// router.HandleFunc("/like-post/{id}", handlers.LikePostHandler).Methods("POST")
	// router.HandleFunc("/like-comment/{id}", handlers.LikeCommentHandler).Methods("POST")
	// router.HandleFunc("/dislike-post/{id}", handlers.DislikePostHandler).Methods("POST")
	// router.HandleFunc("/dislike-comment/{id}", handlers.DislikeCommentHandler).Methods("POST")
	// router.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET")

	//router.HandleFunc("/loggedUser", handlers.LoggedHandler)
}
