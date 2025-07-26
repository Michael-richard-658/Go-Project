package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Michael-richard-658/Go-project/database"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Pass  string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	fmt.Printf("Received user: %+v\n", user)
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	statement, err := database.DB.Prepare(query)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer statement.Close()
	_, err = statement.Exec(user.Name, user.Email, user.Pass)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User created successfully!"))
}
