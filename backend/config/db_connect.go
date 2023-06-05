package config

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func DBConn() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./backend/rltforum.db")
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to the database!")
	return db, nil
}
