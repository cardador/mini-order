package api

import (
	"encoding/json"
	"fmt"
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
	if !eventbus.AddOrder(newOrder) {
		err := fmt.Sprintf("Failed to add order %s", newOrder.ID)
		http.Error(w, err, http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"status": "order_received"})

}

func GetOrder(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	order, exists := store.GetOrder(id)

	if !exists {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(order)

}
