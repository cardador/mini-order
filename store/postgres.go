package store

import (
	"context"
	"database/sql"
	"interview/order/model"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(connStr string) (*PostgresStore, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &PostgresStore{db: db}, nil
}

func (ps *PostgresStore) SaveOrder(ctx context.Context, o model.Order) error {
	query := `INSERT INTO orders (id, item, amount) VALUES ($1, $2, $3)`
	_, err := ps.db.ExecContext(ctx, query, o.ID, o.Item, o.Amount)
	return err
}

func (ps *PostgresStore) GetOrder(ctx context.Context, id string) (model.Order, error) {
	var o model.Order
	query := `SELECT id, item, amount FROM orders WHERE ID = $1`
	err := ps.db.QueryRowContext(ctx, query, id).Scan(&o.ID, &o.Item, &o.Amount)
	return o, err
}
