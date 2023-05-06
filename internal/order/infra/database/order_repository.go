package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joseMarciano/intensive-golang/internal/order/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	statement, err := r.Db.Prepare(
		"INSERT INTO ORDERS (ID, PRICE, TAX, FINAL_PRICE) VALUES (?,?,?,?)",
	)
	if err != nil {
		return err
	}

	_, err = statement.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}

	return nil
}
