package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

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
	fmt.Fprintf(w, "testing backend to frontend")
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
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// handle WebSocket connection here
}

func main() {
	r := mux.NewRouter()
	SetRoutes(r)

	http.Handle("/", r)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8083"},
	})
	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe("localhost:8082", handler))
}

func SetRoutes(router *mux.Router) {
	router.HandleFunc("/socket", handleWebSocket)
	router.HandleFunc("/", IndexHandler)
	// Add more routes as needed
}
