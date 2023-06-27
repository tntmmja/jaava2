package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tntmmja/jaava2/backend/config"
	"github.com/tntmmja/jaava2/backend/data"
	"github.com/tntmmja/jaava2/backend/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Initialize the WebSocket manager
	manager := config.Manager

	// Start the WebSocket manager in a separate goroutine
	go manager.Run()

	// Initialize the database connection
	err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer config.GetDB().Close()

	fmt.Println("Connected main to the database!")

	// Serve static files
	staticDir := "./clientfrontend/static/"
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	// Serve index.html
	r.HandleFunc("/", IndexHandler)

	// Set routes
	SetRoutes(r)

	// Start the HTTP server
	log.Println("Server is running on http://localhost:8082")
	err2 := http.ListenAndServe(":8082", r)
	if err2 != nil {
		log.Fatal("Server failed:", err2)
	}
}

// IndexHandler serves the index.html file
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./clientfrontend/templates/index.html")
}

var SetRoutes = func(router *mux.Router) {
	router.HandleFunc("/", IndexHandler)

	

	router.HandleFunc("/register", data.RegisterHandler).Methods("GET","POST")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./clientfrontend/static/"))))

	router.HandleFunc("/login", handlers.LoginHandler).Methods(("POST"))
	router.HandleFunc("/create-post", handlers.CreatePostHandler).Methods("POST")
	router.HandleFunc("/post/{id}", handlers.PostHandler).Methods("GET")
	router.HandleFunc("/create-comment", handlers.CreateCommentHandler).Methods("POST")
	router.HandleFunc("/comments/{id}", handlers.GetCommentsByPostIDHandler).Methods("GET")
	router.HandleFunc("/comment/{id}", handlers.GetCommentHandler).Methods("GET")
	router.HandleFunc("/send-message", handlers.SendMessageHandler).Methods("POST")
	router.HandleFunc("/messages/{senderID}/{receiverID}", handlers.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/messages/{senderID}/{receiverID}/more", handlers.GetMoreMessagesHandler).Methods("GET")

	router.HandleFunc("/loggedin", handlers.LoggedInHandler).Methods("GET")
	router.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET")

	//router.HandleFunc("/feed", handlers.FeedHandler).Methods("GET", "POST")
	// router.HandleFunc("/filter", handlers.FilterHandler).Methods("GET") // Add this line for the FilterHandler
}
