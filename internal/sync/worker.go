package sync

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"fedratlas-sync/internal/models"
	"fedratlas-sync/internal/peer"
)

func StartSyncWorker() {
	//Actually This function need to be highly change
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

		//Skip if already sent or Failed
		if activity.Status == models.StatusSent || activity.Status == models.StatusFailed {
			continue
		}

		log.Println("Processing activity:", activity.ID)

		success := true

		for _, p := range peers {
			err := sendToPeer(p.BaseURL, activity)
			if err != nil {

				log.Println("Failed sending to:", p.BaseURL)

				success = false
				outbox[i].RetryCount++

				// max retry limit
				if outbox[i].RetryCount >= 3 {
					outbox[i].Status = models.StatusFailed
				}
				log.Println("Activity Failed for", p.ID, "")
				continue
			}
		}

		if success {
			outbox[i].Status = models.StatusSent
			log.Println("Activity synced:", activity.ID)
		}
	}
}

func sendToPeer(baseURL string, activity models.Activity) error {
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
