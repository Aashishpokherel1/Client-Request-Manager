package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"client-request-manager/micro_server" // Import the micro server package

	"github.com/joho/godotenv" // Import the godotenv package
)

// RootHandler handles requests to the root "/"
func RootHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with a simple message indicating this is the main server
	fmt.Fprintf(w, "This is the Main Power Server")
}

// ForwardRequestToReadServer handles GET requests and forwards them to the Read Server
func ForwardRequestToReadServer(w http.ResponseWriter, r *http.Request) {
	// Check the HTTP method type (GET, POST, etc.)
	switch r.Method {
	case http.MethodGet:
		// If the request is a GET method, forward it to the Read Server
		resp, err := http.Get("http://localhost:8083/get") // Read Server's URL
		if err != nil {
			http.Error(w, "Error forwarding request to Read Server", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Respond to the client indicating the request was forwarded
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "GET Request forwarded to Read Server. Status: %d", resp.StatusCode)
		fmt.Fprintf(w, "with url: %d", resp.StatusCode)

	case http.MethodPost:
		// If the request is a POST method, handle it differently
		// You could add some business logic here if needed
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "POST Request received. No further action taken. no post server")

	case http.MethodPut:
		// If the request is a POST method, handle it differently
		// You could add some business logic here if needed
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Put Request received. No further action taken. no put server")
	case http.MethodDelete:
		// If the request is a POST method, handle it differently
		// You could add some business logic here if needed
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "POST Request received. No further action taken.no delete server")

	default:
		// For other methods (PUT, DELETE, etc.), send an error
		http.Error(w, fmt.Sprintf("Method %s not allowed", r.Method), http.StatusMethodNotAllowed)
	}
}

// StartMainServer initializes the routes for the main server
func StartMainServer(done chan bool) {
	// Handle requests for root ("/") and "/power"
	http.HandleFunc("/power", ForwardRequestToReadServer)

	// Handle GET requests and forward to Read Server

	// Get the port from the environment variable, default to ":8082" if not set
	port := os.Getenv("MAIN_SERVER_PORT")
	if port == "" {
		port = ":8082" // default port
	}

	fmt.Printf("Main Power Server started at http://localhost%s\n", port)

	// Start the main server and listen for requests
	go func() {
		log.Fatal(http.ListenAndServe(port, nil))
	}()

	// Notify that the main server has started successfully
	done <- true
}

// Entry point for the application
func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a channel to synchronize the start of the main server
	done := make(chan bool)

	// Start the main server in a goroutine
	go StartMainServer(done)

	// Wait until the main server is up and running
	<-done

	// Now, after the main server is up, start the micro server
	go micro_server.StartReadServer()

	// Keep the main thread running to keep both servers alive
	select {} // Block indefinitely to keep both servers running
}
