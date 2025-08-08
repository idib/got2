package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	got1_client "github.com/idib/got1/pkg/client"

	"github.com/idib/got2/pkg/client"
)

// Server represents the HTTP server
type Server struct {
	addr string
	c    *got1_client.Client
}

// NewServer creates a new server instance
func NewServer(addr string) *Server {

	return &Server{
		addr: addr,
	}
}

// Start initializes and starts the HTTP server
func (s *Server) Start() error {
	mux := http.NewServeMux()

	// Register the ping handler
	mux.HandleFunc("/ping", s.handlePing)
	mux.HandleFunc("/handleChekH", s.handleChekH)

	log.Printf("Starting server on %s", s.addr)
	return http.ListenAndServe(s.addr, mux)
}

// handlePing handles the /ping endpoint
func (s *Server) handlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

// handlePing handles the /ping endpoint
func (s *Server) handleChekH(w http.ResponseWriter, r *http.Request) {

	ss := client.HandleChekHRequest{}

	_ = json.NewDecoder(r.Body).Decode(&ss)

	println(ss)

	fmt.Println("Sending ping request to server...")

	c := got1_client.NewClient("http://localhost:8080")

	// Try to ping the server
	response, err := c.Ping()
	if err != nil {
		log.Fatalf("Failed to ping server: %v", err)
	}

	fmt.Printf("Server response: %s\n", response)

	// Example of how to use the client in a loop
	for i := 1; i <= 3; i++ {
		fmt.Printf("Ping attempt %d: ", i)

		response, err := c.Ping()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Success: %s\n", response)
		}

		time.Sleep(1 * time.Second)
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
