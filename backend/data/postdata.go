// data/postdata.go

package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/tntmmja/jaava2/backend/config"
)

// Post represents a post
type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	
}

// CreatePost creates a new post in the database
func CreatePost(post *Post) error {
	db := config.GetDB()
	if db == nil {
		return errors.New("Failed to get database connection")
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO posts(user_id, title, text, created_at) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(post.UserID, post.Title, post.Text, time.Now())
	if err != nil {
		return err
	}

	postID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	post.ID = int(postID)

	return nil
}

// GetPosts retrieves the posts from the database
func GetPosts() ([]Post, error) {
	db := config.GetDB()
	if db == nil {
		return nil, errors.New("Failed to get database connection")
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, user_id, title, text, created_at FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Text, &post.CreatedAt)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// GetPost retrieves a post from the database by its ID
func GetPost(postID string) (*Post, error) {
	db := config.GetDB()
	if db == nil {
		return nil, errors.New("Failed to get database connection")
	}
	defer db.Close()

	var post Post
	err := db.QueryRow("SELECT id, user_id, title, text, created_at FROM posts WHERE id = ?", postID).Scan(&post.ID, &post.UserID, &post.Title, &post.Text, &post.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Post not found
		}
		return nil, err
	}

	return &post, nil
}

// GetUserIDBySessionID retrieves the user ID associated with a session ID from the database
func GetUserIDBySessionID(sessionID string) (int, error) {
	db := config.GetDB()
	if db == nil {
		return 0, errors.New("Failed to get database connection")
	}
	defer db.Close()

	var userID int
	err := db.QueryRow("SELECT id FROM user WHERE sessionID = ?", sessionID).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil // Session ID not found
		}
		return 0, err
	}

	return userID, nil
}
