package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"fedratlas-sync/internal/models"
	"fedratlas-sync/internal/sync"
)

func CreateActivity(w http.ResponseWriter, r *http.Request) {
	var a models.Activity

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a.ID = uuid.New().String()
	a.Timestamp = time.Now().Unix()

	sync.AddToOutbox(a)

	json.NewEncoder(w).Encode(a)
}
