package handlers

import (
	"net/http"
	"strconv"

	"github.com/Michael-richard-658/Go-project/database"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}
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
	query := "delete from users where id=?"
	result, err := database.DB.Exec(query, userId)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
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
	w.Write([]byte("User deleted :( "))
}
