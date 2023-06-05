package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tntmmja/jaava2/backend/config"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int
	Nickname  string `json:"nickname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	SessionID int
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Submitting registration")
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read and log the request body
	// can delete from 1 here
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Printf("Request body: %s", body)
	// to 1 here delete

	// Decode the JSON data
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	log.Println("Decoded user:", user)

	// Check if nickname is already taken
	if isNicknameTaken(user.Nickname) {
		http.Error(w, "Nickname is already taken", http.StatusBadRequest)
		return
	}

	// Check if email is already taken
	if isEmailTaken(user.Email) {
		http.Error(w, "Email is already taken", http.StatusBadRequest)
		return
	}

	// Generate password hash
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Perform database operations to store user information
	db, err := config.DBConn()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO user (nickname, age, gender, firstName, lastName, email, password) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Nickname, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, hash)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send a response indicating successful registration
	response := map[string]interface{}{
		"message": "Registration successful",
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

// Check if the nickname is already taken
func isNicknameTaken(nickname string) bool {
	// Check the following code for database logic

	db, err := config.DBConn()
	if err != nil {
		log.Println(err)
		return true
	}
	defer db.Close()
	// Perform the query and check if the nickname exists

	row := db.QueryRow("SELECT COUNT(*) FROM user WHERE nickname = ?", nickname)
	var count int
	err = row.Scan(&count)
	if err != nil {
		log.Println(err)
		return true
	}
	return count > 0
	// Always return false (nickname is not taken)
	return false
}

// Check if the email is already taken
func isEmailTaken(email string) bool {
	// Check the following code for database logic

	db, err := config.DBConn()
	if err != nil {
		log.Println(err)
		return true
	}
	defer db.Close()
	// Perform the query and check if the email exists
	row := db.QueryRow("SELECT COUNT(*) FROM user WHERE email = ?", email)
	var count int
	err = row.Scan(&count)
	if err != nil {
		log.Println(err)
		return true
	}
	return count > 0
	return false
}
