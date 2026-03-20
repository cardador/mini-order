package store

import (
	"context"
	"interview/order/model"
)

type Repository interface {
	SaveOrder(ctx context.Context, o model.Order) error
	GetOrder(ctx context.Context, id string) (model.Order, error)
}
