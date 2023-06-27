package config

import (
	"log"

	"github.com/gorilla/websocket"
)

// WebSocketManager manages WebSocket connections
type WebSocketManager struct {
	Connections map[*websocket.Conn]struct{}
	Broadcast   chan []byte
}

// NewWebSocketManager creates a new WebSocketManager instance
func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		Connections: make(map[*websocket.Conn]struct{}),
		Broadcast:   make(chan []byte),
	}
}

// AddConnection adds a WebSocket connection to the manager
func (m *WebSocketManager) AddConnection(conn *websocket.Conn) {
	m.Connections[conn] = struct{}{}
}

// RemoveConnection removes a WebSocket connection from the manager
func (m *WebSocketManager) RemoveConnection(conn *websocket.Conn) {
	delete(m.Connections, conn)
}

// BroadcastMessage broadcasts a message to all WebSocket connections
func (m *WebSocketManager) BroadcastMessage(message []byte) {
	for conn := range m.Connections {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("WebSocket write failed:", err)
			delete(m.Connections, conn)
			conn.Close()
		}
	}
}

// Run starts the WebSocket manager
func (m *WebSocketManager) Run() {
	for {
		select {
		case message := <-m.Broadcast:
			m.BroadcastMessage(message)
		}
	}
}


//without the Manager variable won't be able to access
// the WebSocketManager instance from other parts of the code.
// Manager holds an instance of the WebSocketManager
var Manager = NewWebSocketManager()






// this part is now handled in LoggedInHandler

// // WebSocketUpgrader is the WebSocket upgrader configuration
// var WebSocketUpgrader = websocket.Upgrader{
// 	// ReadBufferSize:  1024,
// 	// WriteBufferSize: 1024,
// 	// CheckOrigin: func(r *http.Request) bool {
// 	// 	return true
// 	// },
// }

// // HandleWebSocket upgrades the HTTP connection to a WebSocket connection and manages it
// func HandleWebSocket(manager *WebSocketManager, w http.ResponseWriter, r *http.Request) {
// 	conn, err := WebSocketUpgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println("WebSocket upgrade failed:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	defer conn.Close()

// 	// Add the WebSocket connection to the manager
// 	manager.AddConnection(conn)

// 	// Handle incoming WebSocket messages
// 	for {
// 		_, message, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println("WebSocket read failed:", err)
// 			break
// 		}

// 		log.Println("Received message:", string(message))

// 		response := []byte("Received your message")
// 		err = conn.WriteMessage(websocket.TextMessage, response)
// 		if err != nil {
// 			log.Println("WebSocket write failed:", err)
// 			break
// 		}
// 	}

// 	// Remove the WebSocket connection from the manager
// 	manager.RemoveConnection(conn)
// }
