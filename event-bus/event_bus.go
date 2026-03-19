package eventbus

import (
	"fmt"
	store "interview/order/store"
	"time"
)

var eventbus = make(chan store.Order, 100)

func ProcessOrder() {
	for order := range eventbus {
		fmt.Printf("[EVENT BUS] Processing payment for order %s...\n", order.ID)
		time.Sleep(2 * time.Second) // heavy work simulation
		fmt.Printf("[EVENT BUS] Order %s complete!\n", order.ID)
	}
}

func AddOrder(o store.Order) {
	eventbus <- o
}
