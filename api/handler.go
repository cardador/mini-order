package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	eventbus "interview/order/event-bus"
	"interview/order/model"
	"interview/order/store"
	"net/http"

	"github.com/lib/pq"
)

func HandleOrder(repo store.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}
		var newOrder model.Order
		if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err := repo.SaveOrder(r.Context(), newOrder)
		if err != nil {
			if pgErr, ok := err.(*pq.Error); ok {
				if pgErr.Code == "23505" {
					http.Error(w, "Order ID already exists", http.StatusConflict)
					return
				}
			}

			http.Error(w, "Failed to save to DB", http.StatusInternalServerError)
			return
		}
		if !eventbus.AddOrder(newOrder) {
			err := fmt.Sprintf("Failed to add order %s", newOrder.ID)
			http.Error(w, err, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{"status": "processing"})
	}
}

func GetOrder(repo store.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
			return
		}

		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}

		order, err := repo.GetOrder(r.Context(), id)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Order not found", http.StatusNotFound)
				return
			}
			msg := fmt.Sprintf("Failed to fetch id: %s from db", id)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(order)
	}
}
