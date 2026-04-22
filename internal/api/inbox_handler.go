package api

import (
	"encoding/json"
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

	sync.AddToInbox(a)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Activity received"))
}
