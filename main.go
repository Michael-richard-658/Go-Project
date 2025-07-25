package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Michael-richard-658/Go-project/handlers"
)

func main() {
	fmt.Println("Server is running on http://localhost:8000")

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/createuser", handlers.CreateUser)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
