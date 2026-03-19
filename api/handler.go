package api

import (
	"encoding/json"
	eventbus "interview/order/event-bus"
	"interview/order/store"
	"net/http"
)

func HandleOrder(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var newOrder store.Order
	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store.SaveOrder(newOrder)
	eventbus.AddOrder(newOrder)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"status": "order_received"})

}
