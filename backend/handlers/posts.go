package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tntmmja/jaava2/backend/data"
)

// CreatePostHandler handles the creation of a new post
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve the post data from the request body
	var post data.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Set the user ID for the post
	sessionID := findSessionCookie(r)
	if sessionID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID, err := data.GetUserIDBySessionID(sessionID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	post.UserID = userID

	// Create the post in the database
	err = data.CreatePost(&post)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return a success response
	response := map[string]interface{}{
		"message": "Post created successfully",
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
}

// PostHandler handles the retrieval of a specific post
func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve the post ID from the URL parameters
	postID := mux.Vars(r)["id"]

	// Get the post from the database
	post, err := data.GetPost(postID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if post == nil {
		http.NotFound(w, r)
		return
	}

	// Return the post as JSON response
	jsonResponse, err := json.Marshal(post)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// FeedHandler retrieves the posts and categories to display in the feed
func FeedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("feedhandler laks kaima")
	switch r.Method {
	case http.MethodGet:
		// Check if the user is logged in
		sessionID := findSessionCookie(r)
		if sessionID == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Perform additional authorization checks if necessary
		// For example, validate the session ID against the database or session store,
		// also check user roles or permissions to determine access rights

		// Retrieve the posts from the database
		posts, err := data.GetPosts()
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Create a response map with the posts data
		response := map[string]interface{}{
			"posts": posts,
		}

		// Convert the response map to JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set the response headers and write the JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)

	case http.MethodPost:
		// Handle POST request to create a new post
		var post data.Post
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			log.Println(err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Set the user ID for the post
		sessionID := findSessionCookie(r)
		if sessionID == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userID, err := data.GetUserIDBySessionID(sessionID)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		post.UserID = userID

		// Create the post in the database
		err = data.CreatePost(&post)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Return a success response
		response := map[string]interface{}{
			"message": "Post created successfully",
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

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

//findSessionCookie is in handle_logged.go now
// func findSessionCookie(r *http.Request) string {
// 	cookie, err := r.Cookie("mycookie")
// 	if err != nil {
// 		return ""
// 	}
// 	return cookie.Value
// }
