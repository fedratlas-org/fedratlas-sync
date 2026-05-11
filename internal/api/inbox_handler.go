package api

import (
	"encoding/json"
	"log"
	"net/http"

	"fedratlas-sync/internal/models"
	"fedratlas-sync/internal/sync"
)

func ReceiveActivity(w http.ResponseWriter, r *http.Request) {
	var a models.Activity

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	added := sync.AddToInbox(a)

	if !added {
		log.Println("Duplicate activity skipped:", a.ID)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Duplicate skipped"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Activity received"))
	log.Println("Activity received", a.ID)
}
