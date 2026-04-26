package sync

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"fedratlas-sync/internal/peer"
)

func StartSyncWorker() {
	go func() {
		for {
			processOutbox()
			time.Sleep(5 * time.Second)
			//sleep time need to be increase, 5 seconds only for demonstration
		}
	}()
}

func processOutbox() {
	activities := GetOutbox()
	peers := peer.GetPeers()

	for i, activity := range activities {

		if activity.Sent {
			continue
		}

		success := true

		for _, p := range peers {
			err := sendToPeer(p.BaseURL, activity)
			if err != nil {
				success = false
			}
		}

		if success {
			outbox[i].Sent = true
		}
	}
}

func sendToPeer(baseURL string, activity interface{}) error {
	url := baseURL + "/inbox"

	body, err := json.Marshal(activity)
	if err != nil {
		log.Println("Failed to marshal:", err)
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("Failed to send to peer:", err)
		return err
	}
	defer resp.Body.Close()

	log.Println("Sent activity to", baseURL)

	return nil
}
