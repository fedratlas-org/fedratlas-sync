package service

import (
	"errors"
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
		Status:     models.StatusPending,
		RetryCount: 0,
	}

	//Send to outbox
	sync.AddToOutbox(activity)

	return activity, nil
}
