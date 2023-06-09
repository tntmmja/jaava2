// data/postdata.go

package data

import (
	"database/sql"
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
	Likes     *int      `json:"like"`
	Dislikes  *int      `json:"dislike"`
}

// CreatePost creates a new post in the database
func CreatePost(post *Post) error {
	db, err := config.DBConn()
	if err != nil {
		return err
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
	db, err := config.DBConn()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, user_id, title, text, created_at, like, dislike FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Text, &post.CreatedAt, &post.Likes, &post.Dislikes)
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
	db, err := config.DBConn()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var post Post
	err = db.QueryRow("SELECT id, user_id, title, text, created_at, like, dislike FROM posts WHERE id = ?", postID).Scan(&post.ID, &post.UserID, &post.Title, &post.Text, &post.CreatedAt, &post.Likes, &post.Dislikes)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Post not found
		}
		return nil, err
	}

	return &post, nil
}
