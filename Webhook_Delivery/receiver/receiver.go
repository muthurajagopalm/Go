package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook-receiver", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		fmt.Println("Received webhook payload:", string(body))
		w.WriteHeader(http.StatusOK)
	})

	log.Println("Receiver listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
