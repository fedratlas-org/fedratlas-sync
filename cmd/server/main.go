package main

import (
	"fedratlas-sync/internal/api"
	"fedratlas-sync/internal/sync"
	"log"
	"net/http"
	"os"

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

	sync.StartSyncWorker()

	var port string

	port = os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	
	log.Println("Server running on :", port)

	err := http.ListenAndServe((":" + port), r)
	if err != nil {
		//This need to change (log.Fatal)
		log.Fatal(err)
	}
}
