package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/tntmmja/jaava2/config"
	"github.com/tntmmja/jaava2/data"
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
	if r.Method == "POST" {
		db := config.GetDB()
		var login Login
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}

		json.Unmarshal([]byte(b), &login)

		fmt.Printf("%s, %s\n", login.Username, login.Password)

		if strings.Trim(login.Username, " ") == "" || strings.Trim(login.Password, " ") == "" {
			fmt.Println("Parameter's can't be empty")
			http.Error(w, "Parameters can't be empty", http.StatusBadRequest)
			return
		}

		var checkUser *sql.Rows
		var err error

		if strings.Contains(login.Username, "@") {
			checkUser, err = db.Query("SELECT id, password, nickname, email FROM user WHERE email=?", login.Username)
		} else {
			checkUser, err = db.Query("SELECT id, password, nickname, email FROM user WHERE nickname=?", login.Username)
		}

		if err != nil {
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
			w.Write([]byte{})
			fmt.Println("suunaloggeduser")
			w.WriteHeader(http.StatusOK)

			var resp = MyResponse{
				Success:   true,
				SessionID: sessionID,
			}

			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			w.Write(jsonResp)
		}
	} else if r.Method == "GET" {
		// Handle GET request
	}
}

type MyResponse struct {
	Success   bool   `json:"success"`
	SessionID string `json:"sessionID"`
}
