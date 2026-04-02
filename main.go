package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status string `json:"status"`
}

// MessageResponse represents the greeting response
type MessageResponse struct {
	Message string `json:"message"`
}

// VersionResponse represents the version response
type VersionResponse struct {
	Version string `json:"version"`
}

// healthHandler handles GET /health requests
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := HealthResponse{Status: "healthy"}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding health response: %v", err)
	}
}

// rootHandler handles GET / requests
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := MessageResponse{Message: "Hello from DevSecOps pipeline"}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding message response: %v", err)
	}
}

// versionHandler handles GET /version requests
func versionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := VersionResponse{Version: "1.0.0"}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding version response: %v", err)
	}
}

// getPort reads the APP_PORT environment variable and returns the port
// Defaults to 8080 if not set
func getPort() string {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Validate that port is a valid number
	if _, err := strconv.Atoi(port); err != nil {
		log.Printf("Invalid APP_PORT value: %s, using default 8080", port)
		return "8080"
	}

	return port
}

func main() {
	port := getPort()

	// Register handlers
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/version", versionHandler)

	// Start server
	address := fmt.Sprintf(":%s", port)
	log.Printf("Starting server on %s", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
