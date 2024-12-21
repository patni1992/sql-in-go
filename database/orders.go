package database

import (
	"database/sql"
)

type OrderRepository struct {
	Db *sql.DB
}

type Order struct {
	Id      int
	Product string
	Amount  int
}

func (r *OrderRepository) CreateTable() error {
	_, err := r.Db.Exec(`CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		product TEXT,
		amount INTEGER
	)`)

	return err
}

func (r *OrderRepository) Insert(order Order) error {
	_, err := r.Db.Exec("INSERT INTO orders (product, amount) VALUES (?, ?)", order.Product, order.Amount)
	return err
}

func (r *OrderRepository) GetAll() ([]Order, error) {
	rows, err := r.Db.Query("SELECT * from orders")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []Order

	for rows.Next() {
		var order Order
		err := rows.Scan(&order.Id, &order.Product, &order.Amount)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepository) GetById(id int) (Order, error) {
	var order Order
	err := r.Db.QueryRow("SELECT * FROM orders WHERE id = ?", id).Scan(&order.Id, &order.Product, &order.Amount)
	if err != nil {
		return Order{}, err
	}

	return order, nil
}

func (r *OrderRepository) Update(order Order) error {
	_, err := r.Db.Exec("UPDATE orders SET product = ?, amount = ? WHERE id = ?", order.Product, order.Amount, order.Id)
	return err
}

func (r *OrderRepository) Delete(id int) error {
	_, err := r.Db.Exec("DELETE FROM orders WHERE id = ?", id)
	return err
}

// Working with context & timeout
// func (r *OrderRepository) test() error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	_, err := r.Db.QueryContext(ctx, "SELECT * from orders")

// 	return err
// }
