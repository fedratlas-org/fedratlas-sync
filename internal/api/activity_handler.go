package api

import (
	"encoding/json"
	"net/http"

	"fedratlas-sync/internal/models"
	"fedratlas-sync/internal/service"
)

func CreateActivity(w http.ResponseWriter, r *http.Request) {
	/*Responsibilities of CreateActivity:
	-Taking input (from API request / internal call)
	-Validating it
	-Structuring it into an Activity object
	-Storing it (DB or memory)
	-Optionally distributing it to peers
	So instead of raw data floating around, everything becomes a standardized activity record*/

	//Method check method should be POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//Content-Type check (application/json)
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	var input models.Activity

	//Decode request from http request
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	//temporary Go object created from request body (a temporary decoded struct, Activity isn't officially created)

	// below 3 added to CreateService func in service(Validation, Creation of Object - Adding metaData, Adding to OutBox)

	//Basic validation check whether required field were missing
	/*if a.Type == "" || a.Actor == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	//Generate metadata
	a.ID = uuid.New().String()
	a.Timestamp = time.Now().Unix()

	//Add to outbox
	sync.AddToOutbox(a)
	*/

	activity, err := service.CreateActivity(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Response to w
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(activity)
}
