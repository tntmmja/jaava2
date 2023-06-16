package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tntmmja/jaava2/backend/config"
	"github.com/tntmmja/jaava2/backend/data"

	"github.com/tntmmja/jaava2/backend/handlers"
	
	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/mux"
	//
)

func main() {
	r := mux.NewRouter()
	SetRoutes(r)

	// Create a new WebSocketManager instance
	manager := config.NewWebSocketManager()

	// Start the WebSocket manager
	go manager.Run()

	// Initialize the database connection
	err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer config.GetDB().Close()

	fmt.Println("Connected main to the database!")

	// // Initialize and start the WebSocket manager
	// WebSocketManager := config.NewWebSocketManager()
	// go WebSocketManager.Run()

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

	router.HandleFunc("/register", data.RegisterHandler).Methods(("POST"))
	router.HandleFunc("/login", handlers.LoginHandler).Methods(("POST"))
	router.HandleFunc("/create-post", handlers.CreatePostHandler).Methods("POST")
	router.HandleFunc("/post/{id}", handlers.PostHandler).Methods("GET")
	router.HandleFunc("/create-comment", handlers.CreateCommentHandler).Methods("POST")
	router.HandleFunc("/comments/{id}", handlers.GetCommentsByPostIDHandler).Methods("GET")
	router.HandleFunc("/comment/{id}", handlers.GetCommentHandler).Methods("GET")
	router.HandleFunc("/send-message", handlers.SendMessageHandler).Methods("POST")
	router.HandleFunc("/messages/{senderID}/{receiverID}", handlers.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/messages/{senderID}/{receiverID}/more", handlers.GetMoreMessagesHandler).Methods("GET")
	
	//router.HandleFunc("/loggedin", handlers.LoggedInHandler)
	

	//router.HandleFunc("/feed", handlers.FeedHandler).Methods("GET", "POST")
	// router.HandleFunc("/filter", handlers.FilterHandler).Methods("GET") // Add this line for the FilterHandler
}
