package main

import (
	"fedratlas-sync/internal/api"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// Register routes
	/*http.HandleFunc("/health", api.HealthCheck)
	http.HandleFunc("/peers/register", api.RegisterPeer)
	http.HandleFunc("/activities", api.CreateActivity)
	http.HandleFunc("/inbox", api.ReceiveActivity)*/

	r.Get("/health", api.HealthCheck)
	r.Post("/peers/register", api.RegisterPeer)
	r.Post("/activities", api.CreateActivity)
	r.Post("/inbox", api.ReceiveActivity)

	log.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		//This need to change (log.Fatal)
		log.Fatal(err)
	}
}
