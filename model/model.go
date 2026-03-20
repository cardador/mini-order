package model

type Order struct {
	ID     string  `json:"id"`
	Item   string  `json:"item"`
	Amount float64 `json:"amount"`
}
