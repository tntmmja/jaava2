package handlers

import (
	"fmt"
	"net/http"
)

// LoggedInHandler handles requests for logged-in users
func LoggedInHandler(w http.ResponseWriter, r *http.Request) {
	// Handle requests for logged-in users
	// ...
	fmt.Fprintln(w, "Welcome, logged-in user!")
}
