package sync

import "fedratlas-sync/internal/models"

var outbox []models.Activity

func AddToOutbox(a models.Activity) {
	outbox = append(outbox, a)
}

func GetOutbox() []models.Activity {
	return outbox
}

func findActivity(activityID string) *models.Activity {

	for i := range outbox {

		if outbox[i].ID == activityID {
			return &outbox[i]
		}
	}

	return nil
}

func findDelivery(
	activity *models.Activity,
	peerID string,
) *models.PeerDelivery {

	for i := range activity.Deliveries {

		if activity.Deliveries[i].PeerID == peerID {
			return &activity.Deliveries[i]
		}
	}

	return nil
}

func IncrementDeliveryRetry(
	activityID string,
	peerID string,
) {
	activity := findActivity(activityID)

	if activity == nil {
		return
	}

	delivery := findDelivery(activity, peerID)

	if delivery == nil {
		return
	}

	delivery.RetryCount++

	if delivery.RetryCount >= 3 {
		delivery.Status = models.DeliveryFailed
	}
}

func MarkDeliverySent(
	activityID string,
	peerID string,
) {
	activity := findActivity(activityID)

	if activity == nil {
		return
	}

	delivery := findDelivery(activity, peerID)

	if delivery == nil {
		return
	}

	delivery.Status = models.DeliverySent
}

//There is No use of below function we can have the same functionality only with incrementDeliveryRetryCount func
/*func MarkDeliveryFailed(
	activityID string,
	peerID string,
) {
	activity := findActivity(activityID)

	if activity == nil {
		return
	}

	delivery := findDelivery(activity, peerID)

	if delivery == nil {
		return
	}

	delivery.Status = models.DeliveryFailed
}*/

func IsActivityFullyDelivered(activityID string) bool {
	activity := findActivity(activityID)

	if activity == nil {
		return false
	}

	for _, delivery := range activity.Deliveries {
		if delivery.Status != models.DeliverySent {
			return false
		}
	}

	return true
}
