package models

type ActivityType string

const (
	AddFeature    ActivityType = "ADD_FEATURE"
	UpdateFeature ActivityType = "UPDATE_FEATURE"
	DeleteFeature ActivityType = "DELETE_FEATURE"
)

type Activity struct {
	ID        string       `json:"id"`
	Type      ActivityType `json:"type"`
	Timestamp int64        `json:"timestamp"`
	Actor     string       `json:"actor"`
	Object    interface{}  `json:"object"`
	Sent      bool         `json:"sent"` //Check Whether the activity sent to peers or not
}
