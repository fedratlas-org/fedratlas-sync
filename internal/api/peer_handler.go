package api

import (
	"encoding/json"
	"net/http"

	"fedratlas-sync/internal/models"
	"fedratlas-sync/internal/peer"
)

func RegisterPeer(w http.ResponseWriter, r *http.Request) {
	var p models.Peer

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	peer.AddPeer(p)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Peer registered"))
}
