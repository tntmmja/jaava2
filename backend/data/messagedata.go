// data/messagedata.go

package data

import (
	"errors"
	"time"

	"github.com/tntmmja/jaava2/backend/config"
)

// Message represents a message sent between users
type Message struct {
	ID         int       `json:"id"`
	SenderID   int       `json:"sender_id"`
	ReceiverID int       `json:"receiver_id"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"created_at"`
}

// SendMessage sends a new message between users and stores it in the database
func SendMessage(message *Message) error {
	db := config.GetDB()
	if db == nil {
		return errors.New("Failed to get database connection")
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO messages(sender_id, receiver_id, message, created_at) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(message.SenderID, message.ReceiverID, message.Message, time.Now())
	if err != nil {
		return err
	}

	return nil
}

// GetMessagesBetweenUsers retrieves the last 10 messages between two users based on their IDs
func GetMessagesBetweenUsers(senderID, receiverID int) ([]Message, error) {
	db := config.GetDB()
	if db == nil {
		return nil, errors.New("Failed to get database connection")
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, sender_id, receiver_id, message, created_at FROM messages WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?) ORDER BY created_at DESC LIMIT 10", senderID, receiverID, receiverID, senderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message

	for rows.Next() {
		var message Message
		err = rows.Scan(&message.ID, &message.SenderID, &message.ReceiverID, &message.Message, &message.CreatedAt)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Reverse the order of messages to display them from oldest to newest
	for i := len(messages)/2 - 1; i >= 0; i-- {
		opp := len(messages) - 1 - i
		messages[i], messages[opp] = messages[opp], messages[i]
	}

	return messages, nil
}

// GetMoreMessages retrieves additional messages between two users based on an offset and limit, allowing for pagination
func GetMoreMessages(senderID, receiverID int, offset, limit int) ([]Message, error) {
	db := config.GetDB()
	if db == nil {
		return nil, errors.New("Failed to get database connection")
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, sender_id, receiver_id, message, created_at FROM messages WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?) ORDER BY created_at DESC LIMIT ? OFFSET ?", senderID, receiverID, receiverID, senderID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message

	for rows.Next() {
		var message Message
		err = rows.Scan(&message.ID, &message.SenderID, &message.ReceiverID, &message.Message, &message.CreatedAt)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Reverse the order of messages to display them from oldest to newest
	for i := len(messages)/2 - 1; i >= 0; i-- {
		opp := len(messages) - 1 - i
		messages[i], messages[opp] = messages[opp], messages[i]
	}

	return messages, nil
}
