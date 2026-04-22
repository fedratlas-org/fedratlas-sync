package sync

import "fedratlas-sync/internal/models"

var outbox []models.Activity

func AddToOutbox(a models.Activity) {
	outbox = append(outbox, a)
}

func GetOutbox() []models.Activity {
	return outbox
}
