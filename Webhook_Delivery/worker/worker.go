package worker

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/muthurajagopalm/Webhook_Delivery/models"
	"github.com/muthurajagopalm/Webhook_Delivery/queue"
)

func Start(workerCount int) {
	for i := 0; i < workerCount; i++ {
		go func(id int) {
			for event := range queue.WebhookQueue {
				log.Printf("[Worker %d] Sending to %s", id, event.TargetURL)
				err := deliver(event)
				if err != nil && event.RetryCount < 3 {
					event.RetryCount++
					time.Sleep(time.Duration(event.RetryCount) * 2 * time.Second)
					queue.Enqueue(event)
				} else if err != nil {
					log.Printf("[Worker %d] failed permanently: %v", id, err)
				}
			}

		}(i)
	}

}

func deliver(event models.WebhookEvent) error {
	body, _ := json.Marshal(event.Payload)
	req, _ := http.NewRequest("POST", event.TargetURL, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Printf("delivered successfully to %s", event.TargetURL)
		return nil
	}

	return err
}
