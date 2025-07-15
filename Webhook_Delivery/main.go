package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/muthurajagopalm/Webhook_Delivery/models"
	"github.com/muthurajagopalm/Webhook_Delivery/queue"
	"github.com/muthurajagopalm/Webhook_Delivery/worker"
)

func main() {
	// Initialize the webhook queue with a buffer size of 100
	queue.InitQueue(100)

	// Start the worker to process webhook events
	go worker.Start(5)
	// Set up the HTTP server to handle incoming webhook events
	http.HandleFunc("/send-webhook", handleWebhook)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON payload
	var event models.WebhookEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	log.Printf("Received webhook event: %s", event.ID)
	queue.Enqueue(event)
	w.WriteHeader(http.StatusAccepted)
}
