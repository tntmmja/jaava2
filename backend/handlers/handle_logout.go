// handle_logout.go

package handlers

import (
	"net/http"
)

// LogoutHandler handles the /logout route
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the user is logged in
	sessionID := findSessionCookie(r)
	if sessionID == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Clear the session cookie
	clearSessionCookie(w)

	// Redirect to the homepage or login page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func clearSessionCookie(w http.ResponseWriter) {
	// Clear the session cookie in the response
	cookie := &http.Cookie{
		Name:   "mycookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // Set the MaxAge to -1 to delete the cookie
	}
	http.SetCookie(w, cookie)
}
