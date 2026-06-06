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

type DeliveryStatus string

const (
	DeliveryPending DeliveryStatus = "PENDING"
	DeliverySent    DeliveryStatus = "SENT"
	DeliveryFailed  DeliveryStatus = "FAILED"
)

type PeerDelivery struct {
	PeerID     string         `json:"peer_id"`
	Status     DeliveryStatus `json:"status"`
	RetryCount int            `json:"retry_count"`
}

type Activity struct {
	ID         string         `json:"id"`
	Type       ActivityType   `json:"type"`
	Timestamp  int64          `json:"timestamp"`
	Actor      string         `json:"actor"`
	Object     interface{}    `json:"object"`
	Status     ActivityStatus `json:"status"` //Check Whether the activity sent to all peers or not
	RetryCount int            `json:"retryCount"`

	Deliveries []PeerDelivery `json:"deliveries"`
}
