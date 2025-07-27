package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Michael-richard-658/Go-project/database"
	"github.com/Michael-richard-658/Go-project/handlers"
)

func main() {
	database.DbConnection()
	defer database.DB.Close()

	// Create an instance of UserCRUD
	UserCRUD := &handlers.UserCRUD{}

	fmt.Println("Server is running on http://localhost:8000")

	// Use instance methods
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/createuser", UserCRUD.CreateUser)
	http.HandleFunc("/loginuser", UserCRUD.LoginUser)
	http.HandleFunc("/edituser", UserCRUD.EditUser)
	http.HandleFunc("/deleteuser", UserCRUD.DeleteUser)

	// Start the server
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
