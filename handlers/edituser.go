package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Michael-richard-658/Go-project/database"
)

func (u UserCRUD) EditUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get user ID from header
	getId := r.Header.Get("id")
	if getId == "" {
		http.Error(w, "Missing user ID header", http.StatusBadRequest)
		return
	}
	userId, err := strconv.Atoi(getId)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Decode JSON body into user struct
	var updatedUser User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Execute update query
	query := "UPDATE users SET name=?, email=?, password=? WHERE id=?"
	result, err := database.DB.Exec(query, updatedUser.Name, updatedUser.Email, updatedUser.Pass, userId)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Check if any row was affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Could not determine update result", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "User does not exist", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "User with ID %d updated successfully", userId)
}
