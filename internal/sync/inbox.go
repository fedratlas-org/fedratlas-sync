package sync

import "fedratlas-sync/internal/models"

var inbox []models.Activity

func AddToInbox(a models.Activity) {
	inbox = append(inbox, a)
}

func GetInbox() []models.Activity {
	return inbox
}
