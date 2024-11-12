package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Client struct defines the structure of a client
type Client struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// RootHandler handles requests to the root "/"
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Client Request Manager!")
}

// ClientsHandler handles requests to "/clients"
func ClientsHandler(w http.ResponseWriter, r *http.Request) {
	clients := []Client{
		{ID: 1, Name: "Client A", Email: "clienta@example.com"},
		{ID: 2, Name: "Client B", Email: "clientb@example.com"},
	}

	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Send the list of clients as JSON
	if err := json.NewEncoder(w).Encode(clients); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

// StartServer initializes the routes and starts the HTTP server
func StartServer() {
	// Handle requests for root ("/") and "/clients"
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/clients", ClientsHandler)

	// Define the address and port for the server (e.g., :8080)
	port := ":8080"
	fmt.Printf("Server started at http://localhost%s\n", port)

	// Start the server and listen for requests
	log.Fatal(http.ListenAndServe(port, nil))
}

// Entry point for the application
func main() {
	// Start the server
	StartServer()
}
