package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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
	db := config.GetDB()
	if db == nil {
		log.Println("Failed to get database connection")
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
	 // Encode the response as JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
 	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Write the JSON response
	w.Write(jsonResponse)
}

// Check if the nickname is already taken
func isNicknameTaken(nickname string) bool {
	// Check the following code for database logic
	db := config.GetDB()
	if db == nil {
		log.Println("Failed to get database connection")
		return true
	}
	// Perform the query and check if the nickname exists
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM user WHERE nickname = ?", nickname).Scan(&count)
	if err != nil {
		log.Println(err)
		return true
	}

	return count > 0
}

// Check if the email is already taken
func isEmailTaken(email string) bool {
	// Check the following code for database logic

	db := config.GetDB()
	if db == nil {
		log.Println("Failed to get database connection")
		return true
	}
	// Perform the query and check if the email exists
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM user WHERE email = ?", email).Scan(&count)
	if err != nil {
		log.Println(err)
		return true
	}

	return count > 0
}

func renderRegistrationForm(w http.ResponseWriter, errorMessage string) {
	// Read the contents of the registration form HTML template file
	templateFile := "./clientfrontend/templates/register.html"
	templateBytes, err := ioutil.ReadFile(templateFile)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Convert the template bytes to a string
	templateStr := string(templateBytes)

	// Replace the placeholder in the template with the error message
	renderedTemplate := strings.Replace(templateStr, "{{ErrorMessage}}", errorMessage, 1)

	// Write the rendered template to the response
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(renderedTemplate))
}
