package service

import (
	"errors"
	"fedratlas-sync/internal/peer"
	"time"

	"github.com/google/uuid"

	"fedratlas-sync/internal/models"
	"fedratlas-sync/internal/sync"
)

func CreateActivity(input models.Activity) (models.Activity, error) {
	//Validation

	if input.Type == "" {
		return models.Activity{}, errors.New("type is required")
	}
	if input.Actor == "" {
		return models.Activity{}, errors.New("actor is required")
	}

	//Create new activity (input -> New activity & adding metadata)
	activity := models.Activity{
		ID:         uuid.New().String(),
		Type:       input.Type,
		Actor:      input.Actor,
		Object:     input.Object,
		Timestamp:  time.Now().Unix(),
		Status:     models.StatusPending, //Need to be removed when Deliveries Successfully implemented
		RetryCount: 0,                    //Need to be removed when Deliveries Successfully implemented

		Deliveries: buildDeliveries(),
	}

	//Send to outbox
	sync.AddToOutbox(activity)

	return activity, nil
}

// below Function is responsible for adding all the peers to activity.Deliveries when activity get created
func buildDeliveries() []models.PeerDelivery {

	peers := peer.GetPeers()

	deliveries := []models.PeerDelivery{}

	for _, p := range peers {

		deliveries = append(deliveries, models.PeerDelivery{
			PeerID:     p.ID,
			Status:     models.DeliveryPending,
			RetryCount: 0,
		})
	}

	return deliveries
}
