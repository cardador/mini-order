package model

type Order struct {
	ID     string  `json:"id" dynamodbav:"id"`
	Item   string  `json:"item" dynamodbav:"item"`
	Amount float64 `json:"amount" dynamodbav:"amount"`
}
