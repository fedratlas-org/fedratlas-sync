package main

import (
	"fedratlas-sync/internal/api"
	"log"
	"net/http"
)

func main() {
	// Register routes
	http.HandleFunc("/health", api.HealthCheck)
	http.HandleFunc("/peers/register", api.RegisterPeer)
	http.HandleFunc("/activities", api.CreateActivity)
	http.HandleFunc("/inbox", api.ReceiveActivity)

	log.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
