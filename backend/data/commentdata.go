// data/comments.go

package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/tntmmja/jaava2/backend/config"
)

// Comment represents a comment on a post
type Comment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateComment creates a new comment in the database
func CreateComment(comment *Comment) error {
	db := config.GetDB()
	if db == nil {
		return errors.New("Failed to get database connection")
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO comments(user_id, post_id, text, created_at) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(comment.UserID, comment.PostID, comment.Text, time.Now())
	if err != nil {
		return err
	}

	return nil
}

// GetCommentsByPostID retrieves the comments for a specific post from the database


func GetCommentsByPostID(postID string) ([]Comment, error) {
	db := config.GetDB()
	if db == nil {
		return nil, errors.New("Failed to get database connection")
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, user_id, post_id, text, created_at FROM comments WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment

	for rows.Next() {
		var comment Comment
		err = rows.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Text, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

// GetComment retrieves a comment from the database by its ID
func GetComment(commentID string) (*Comment, error) {
	db := config.GetDB()
	if db == nil {
		return nil, errors.New("Failed to get database connection")
	}
	defer db.Close()

	var comment Comment
	err := db.QueryRow("SELECT id, user_id, post_id, text, created_at FROM comments WHERE id = ?", commentID).Scan(
		&comment.ID, &comment.UserID, &comment.PostID, &comment.Text, &comment.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Comment not found
		}
		return nil, err
	}

	return &comment, nil
}
