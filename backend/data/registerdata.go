package data

import (
	//"encoding/json"
	"fmt"
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
	html := `
	<h1>Registration Form</h1>
	<form id="registration-form" method="POST" action="/register">
		<label for="nickname">Nickname:</label>
		<input type="text" id="nickname" name="nickname"><br><br>
		
		<label for="age">Age:</label>
		<input type="number" id="age" name="age"><br><br>
		
		<label for="gender">Gender:</label>
		<input type="text" id="gender" name="gender"><br><br>
		
		<label for="firstName">First Name:</label>
		<input type="text" id="firstName" name="firstName"><br><br>
		
		<label for="lastName">Last Name:</label>
		<input type="text" id="lastName" name="lastName"><br><br>
		
		<label for="email">E-mail:</label>
		<input type="email" id="email" name="email"><br><br>
		
		<label for="password">Password:</label>
		<input type="password" id="password" name="password"><br><br>
		
		<input type="submit" value="Register">
	</form>
	<a href="/">Go to Login</a>
	`
	if r.Method == http.MethodGet {
		// Serve the registration form
		// Construct and return the registration form HTML
		

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(html))
	} else if r.Method == http.MethodPost {
		// Decode the form data
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Retrieve the form values
		nickname := r.Form.Get("nickname")
		age := r.Form.Get("age")
		gender := r.Form.Get("gender")
		firstName := r.Form.Get("firstName")
		lastName := r.Form.Get("lastName")
		email := r.Form.Get("email")
		password := r.Form.Get("password")

		// Check if nickname is already taken
		if isNicknameTaken(nickname) {
			http.Error(w, "Nickname is already taken", http.StatusBadRequest)
			return
		}

		// Check if email is already taken
		if isEmailTaken(email) {
			http.Error(w, "Email is already taken", http.StatusBadRequest)
			return
		}

		// Generate password hash
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
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

		_, err = stmt.Exec(nickname, age, gender, firstName, lastName, email, hash)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// // Send a response indicating successful registration
		// response := map[string]interface{}{
		// 	"message": "Registration successful",
		// }
		// // Encode the response as JSON
		// jsonResponse, err := json.Marshal(response)
		// if err != nil {
		// 	log.Println(err)
		// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// 	return
		// }
		// // Set the response headers
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// // Write the JSON response
		// w.Write(jsonResponse)

		 // Send the success message
		 fmt.Fprint(w, html)
		 fmt.Fprint(w, "<p>Registration successful</p>")

		return

	} else {
		// Return a Method Not Allowed response for other HTTP methods
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
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

// func renderRegistrationForm(w http.ResponseWriter, errorMessage string) {
// 	// Read the contents of the registration form HTML template file
// 	templateFile := "./clientfrontend/templates/register.html"
// 	templateBytes, err := ioutil.ReadFile(templateFile)
// 	if err != nil {
// 		log.Println(err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	// Convert the template bytes to a string
// 	templateStr := string(templateBytes)

// 	// Replace the placeholder in the template with the error message
// 	renderedTemplate := strings.Replace(templateStr, "{{ErrorMessage}}", errorMessage, 1)

// 	// Write the rendered template to the response
// 	w.Header().Set("Content-Type", "text/html")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(renderedTemplate))
// }
