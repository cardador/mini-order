package store

import "sync"

type Order struct {
	ID     string  `json:"id"`
	Item   string  `json:"item"`
	Amount float64 `json:"amount"`
}

var (
	db     = make(map[string]Order)
	dbLock sync.RWMutex
)

func SaveOrder(o Order) {
	dbLock.Lock()
	defer dbLock.Unlock()
	db[o.ID] = o
}

func GetOrder(id string) (Order, bool) {
	dbLock.RLock()
	defer dbLock.RLocker().Unlock()
	val, exists := db[id]
	return val, exists
}
