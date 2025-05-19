package order

import (
	"database/sql"

	"github.com/XohChen/NewGOBackEnd/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	res, err := s.db.Exec("INSERT INTO order (userId, total,status,address,createdAt) VALUES (?,?,?,?,?)",
		order.ID, order.UserID, order.Total, order.Status, order.Address, order.CreatedAt)
	if err != nil {
		return 0, nil
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items (id, orderId, productId, quantity, price, createdAt) VALUES (?,?,?,?,?,?)",
		orderItem.ID, orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price, orderItem.CreatedAt)

	return err
}
