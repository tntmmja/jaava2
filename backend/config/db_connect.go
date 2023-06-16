package config

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitDB initializes the database connection.
// this is a centralized database connection management approach.


func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./backend/rltforum.db")
	if err != nil {
		return err
	}

	fmt.Println("Connected to the database!")

	// Set database connection settings
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	return nil
}

// GetDB returns the active database connection.
func GetDB() *sql.DB {
	return db
}
