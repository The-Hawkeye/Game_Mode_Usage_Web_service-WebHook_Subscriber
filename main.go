package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


type WebhookReceiver struct{}

func (wr *WebhookReceiver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload struct {
		AreaCode string `json:"area_code"`
		Mode     string `json:"mode"`
		Count    int    `json:"count"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Failed to decode payload", http.StatusBadRequest)
		return
	}


	fmt.Printf("Received webhook notification: %+v\n", payload)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.Handle("/webhook", &WebhookReceiver{})
	fmt.Println("Listening for webhook notifications on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
