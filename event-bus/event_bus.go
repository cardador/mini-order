package eventbus

import (
	"fmt"
	"interview/order/model"
	"time"
)

var eventbus = make(chan model.Order, 100)

func ProcessOrder() {
	for order := range eventbus {
		fmt.Printf("[EVENT BUS] Processing payment for order %s...\n", order.ID)
		time.Sleep(2 * time.Second) // heavy work simulation
		fmt.Printf("[EVENT BUS] Order %s complete!\n", order.ID)
	}
}

func AddOrder(o model.Order) bool {
	select {
	case eventbus <- o:
		return true
	default:
		fmt.Printf("[WARN] Event bus full, dropping order %s\n", o.ID)
		return false
	}
}
