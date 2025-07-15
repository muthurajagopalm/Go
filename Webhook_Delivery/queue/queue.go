package queue

import "github.com/muthurajagopalm/Webhook_Delivery/models"

var WebhookQueue chan models.WebhookEvent

func InitQueue(buffer int) {
	WebhookQueue = make(chan models.WebhookEvent, buffer)
}

func Enqueue(event models.WebhookEvent) {
	WebhookQueue <- event
}
