package models

type WebhookEvent struct {
	ID         string `json:"id"`
	Event      string `json:"event"`
	TargetURL  string `json:"target_url"`
	Payload    any    `json:"payload"`
	RetryCount int    `json:"retry_count"`
}
