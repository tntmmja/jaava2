package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/tntmmja/jaava2/backend/config"
)

// verifying the user's session and authentication status before allowing logged user activities
// LoggedInHandler handles requests for logged-in users
func LoggedInHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("loggedin handler")
	// Check if the user is logged in
	sessionID := findSessionCookie(r)
	fmt.Println("loggedin kuuki", sessionID)
	if sessionID == "" {
		fmt.Println("loggedin sessionid pon null")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	fmt.Println("loggedin kuuki2")
	// Perform additional authorization checks if necessary
	// For example, you can validate the session ID against the database or session store
	// You can also check user roles or permissions to determine access rights

	// Handle requests for logged-in users
	// In this example, we simply return a success response indicating the user is logged in
	response := map[string]interface{}{
		"message": "Welcome, logged-in user!",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)




// // Check for custom headers
// customHeader1 := r.Header.Get("Custom-Header-1")
// customHeader2 := r.Header.Get("Custom-Header-2")

// // Validate the custom headers if necessary
// if customHeader1 != "Value1" || customHeader2 != "Value2" {
// 	log.Println("Invalid headers:", customHeader1, customHeader2)
//   http.Error(w, "Invalid headers", http.StatusBadRequest)
//   return
// }



	// Upgrade the connection to a WebSocket connection
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgreider feilis")
		log.Println("WebSocket upgrade failed:", err)
		log.Println("Internal Server Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Add the WebSocket connection to the manager
	config.Manager.AddConnection(conn)

	// Handle incoming WebSocket messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read failed:", err)
			break
		}

		log.Println("Received message:", string(message))

		response := []byte("Received your message")
		err = conn.WriteMessage(websocket.TextMessage, response)
		if err != nil {
			log.Println("WebSocket write failed:", err)
			break
		}
	}

	// Remove the WebSocket connection from the manager
	config.Manager.RemoveConnection(conn)

}

func findSessionCookie(r *http.Request) string {
	cookie, err := r.Cookie("mycookie")
	if err != nil {
		return ""
	}
	return cookie.Value
}
