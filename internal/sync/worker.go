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

type SyncResult struct {
	PeerID  string
	Success bool
	Error   error
}

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

		//Skip if already sent or Failed
		if activity.Status == models.StatusSent || activity.Status == models.StatusFailed {
			continue
		}
		//need to keep this because if the activity sent to all peers there is no use of checking them one by one

		log.Println("Processing activity:", activity.ID)

		results := make(chan SyncResult)
		pendingCount := 0

		for _, delivery := range activity.Deliveries {

			if delivery.Status != models.DeliveryPending {
				continue
			}

			pendingCount++

			//In below func we could only get the peerID as the parameter But It could be a problem in future
			//That's why we didn't do it
			go func(peerID string, peerBaseURL string) {
				err := sendToPeer(peerBaseURL, activity)

				results <- SyncResult{
					PeerID:  peerID,
					Success: err == nil,
					Error:   err,
				}

			}(delivery.PeerID, peers[delivery.PeerID].BaseURL)
		}

		//Below Part is for increment Retry Count and Mark as sent
		//success := true

		for range pendingCount {
			result := <-results

			if !result.Success {

				//success = false

				//outbox[i].RetryCount++
				IncrementDeliveryRetry(activity.ID, result.PeerID)

				log.Println("Failed:", result.PeerID)

				/*if outbox[i].RetryCount >= 3 {
					outbox[i].Status = models.StatusFailed
					log.Println("Activity Failed for", result.PeerID, "")
					continue
				}*/

				continue

			}

			MarkDeliverySent(activity.ID, result.PeerID)
		}

		/*for range peers {

			result := <-results

			if !result.Success {

				success = false

				outbox[i].RetryCount++

				log.Println("Failed:", result.PeerID)

				if outbox[i].RetryCount >= 3 {
					outbox[i].Status = models.StatusFailed

					log.Println("Activity Failed for", result.PeerID, "")
					continue
				}

			}
		}*/

		if IsActivityFullyDelivered(activity.ID) {
			outbox[i].Status = models.StatusSent
			log.Println("Activity synced:", activity.ID)
		}

		//This part need to improve
		activities[i].RetryCount++
		if activities[i].RetryCount >= 5 {
			activities[i].Status = models.StatusFailed
		}
	}
}

func sendToPeer(baseURL string, activity models.Activity) error {
	url := baseURL + "/inbox"

	//marshals models.Activity to json object
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
