package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tntmmja/jaava2/backend/data"
)

// FeedHandler retrieves the posts to display in the feed
func FeedHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve the posts from the database
	posts, err := data.GetPosts()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Convert the posts data to JSON
	jsonData, err := json.Marshal(posts)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the response headers and write the JSON data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
