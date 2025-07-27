package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Michael-richard-658/Go-project/database"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u UserCRUD) LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Use QueryRow to fetch a single user
	query := "SELECT email FROM users WHERE email = ? AND password = ?"
	row := database.DB.QueryRow(query, credentials.Email, credentials.Password)

	var email string
	err = row.Scan(&email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful!"))
}
