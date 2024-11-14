package micro_server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// ClientsHandler handles GET requests to /get (micro server)
func ClientsHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with a simple message indicating this is the micro server
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "This is the Micro Server at /get")
}

// StartReadServer starts the read-only server for handling GET requests
func StartReadServer() {
	// Handle GET requests to /get (micro server)
	http.HandleFunc("/get", ClientsHandler)

	// Get the port from the environment variable, default to ":8083" if not set
	port := os.Getenv("MICRO_SERVER_PORT")
	if port == "" {
		port = ":8083" // default port
	}

	fmt.Printf("Micro Server started at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
