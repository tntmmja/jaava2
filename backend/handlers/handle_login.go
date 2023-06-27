package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/tntmmja/jaava2/backend/config"
	"github.com/tntmmja/jaava2/backend/data"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	ID        int
	Username  string `json:"username"`
	Password  string `json:"password"`
	SessionID string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginhandler")
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read and log the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Printf("Request body: %s", body)

	// Decode the JSON data
	var login Login
	err = json.Unmarshal(body, &login)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	fmt.Printf("%s, %s\n", login.Username, login.Password)

	if strings.Trim(login.Username, " ") == "" || strings.Trim(login.Password, " ") == "" {
		fmt.Println("Parameter's can't be empty")
		http.Error(w, "Parameters can't be empty", http.StatusBadRequest)
		return
	}


	

// Perform database operations to validate login credentials
	db := config.GetDB() // Obtain the database connection from the middleware
	if db == nil {
		log.Println("Failed to get database connection")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var checkUser *sql.Rows
	var err2 error

	if strings.Contains(login.Username, "@") {
		checkUser, err2 = db.Query("SELECT id, password, nickname, email FROM user WHERE email=?", login.Username)
	} else {
		checkUser, err2 = db.Query("SELECT id, password, nickname, email FROM user WHERE nickname=?", login.Username)
	}

	if err2 != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer checkUser.Close()
	user := &data.User{}
	for checkUser.Next() {
		var id int
		var password, nickName, email string
		err = checkUser.Scan(&id, &password, &nickName, &email)

		log.Println("------------------------", id, password, nickName, email)
		if err != nil {
			fmt.Println("------------", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		user.ID = id
		user.Nickname = nickName
		user.Email = email
		user.Password = password
	}

	if user.ID == 0 {
		http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
		return
	}

	var resp MyResponse
	


	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		fmt.Println("Invalid password")
		http.Error(w, "Invalid password", http.StatusUnauthorized)
	} else {


		sessionID := uuid.New().String()
		fmt.Println("loginsessionid1", sessionID)


		upt, err := db.Prepare("UPDATE user SET sessionID = ? WHERE id = ?")

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer upt.Close()
		_, err = upt.Exec(sessionID, user.ID)
		login.SessionID = sessionID
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
			

		w.Header().Add("Set-Cookie", "mycookie="+sessionID+"; Max-Age=300")

		fmt.Println("login handleri suunaloggeduser")

		fmt.Println("Redirecting to the loggedin page")
		resp = MyResponse{
			Success:   true,
			SessionID: sessionID,
		}

		//http.Redirect(w, r, "/loggedin", http.StatusSeeOther)
		//return

		
	}

	// var resp = MyResponse{
	// 	Success:   true,
	// 	SessionID: login.SessionID, // Update with the correct variable name
	// }

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	// Write the response body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)



}

type MyResponse struct {
	Success   bool   `json:"success"`
	SessionID string `json:"sessionID"`
}
