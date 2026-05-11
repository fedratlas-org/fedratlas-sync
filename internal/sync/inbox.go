package sync

import "fedratlas-sync/internal/models"

var inbox []models.Activity
var processedActivities = make(map[string]bool)

func AddToInbox(a models.Activity) bool {
	// already processed
	if processedActivities[a.ID] {
		return false
	}

	inbox = append(inbox, a)
	processedActivities[a.ID] = true

	return true
}

func GetInbox() []models.Activity {
	return inbox
}
