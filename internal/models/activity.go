package models

type ActivityType string

const (
	AddFeature    ActivityType = "ADD_FEATURE"
	UpdateFeature ActivityType = "UPDATE_FEATURE"
	DeleteFeature ActivityType = "DELETE_FEATURE"
)

type ActivityStatus string

const (
	StatusPending ActivityStatus = "PENDING"
	StatusSent    ActivityStatus = "SENT"
	StatusFailed  ActivityStatus = "FAILED"
)

type Activity struct {
	ID         string         `json:"id"`
	Type       ActivityType   `json:"type"`
	Timestamp  int64          `json:"timestamp"`
	Actor      string         `json:"actor"`
	Object     interface{}    `json:"object"`
	Sent       bool           `json:"sent"`   //This need to be removed after fully implemented the Status
	Status     ActivityStatus `json:"status"` //Check Whether the activity sent to peers or not(previously This was done by sent bool)
	RetryCount int            `json:"retryCount"`
}
