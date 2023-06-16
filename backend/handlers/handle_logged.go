package handlers

import (
	"fmt"
	"net/http"
)
// verifying the user's session and authentication status before allowing logged user activities
// LoggedInHandler handles requests for logged-in users
func LoggedInHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the user is logged in
	sessionID := findSessionCookie(r)
	if sessionID == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Perform additional authorization checks if necessary
	// For example, you can validate the session ID against the database or session store
	// You can also check user roles or permissions to determine access rights

	// Handle requests for logged-in users
	// ...

	fmt.Fprintln(w, "Welcome, logged-in user!")
}




func findSessionCookie(r *http.Request) string {
	cookie, err := r.Cookie("mycookie")
	if err != nil {
		return ""
	}
	return cookie.Value
}
