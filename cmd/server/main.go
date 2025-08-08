package main

import (
	"flag"
	"log"

	"github.com/idib/got2/internal/server"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP server address")
	flag.Parse()

	srv := server.NewServer(*addr)
	log.Printf("Server starting on %s", *addr)

	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
