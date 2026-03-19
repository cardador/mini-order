package main

import (
	"fmt"
	"interview/order/api"
	eventbus "interview/order/event-bus"
	"net/http"
)

func main() {

	go eventbus.ProcessOrder()

	http.HandleFunc("/order", api.HandleOrder)

	fmt.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server failed:", err)
	}
}
