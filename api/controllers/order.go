package controllers

import (
	"fmt"
	"log"
)

type Order struct {
	Id          int
	User_id     string
	Product_id  int
	Quantity    int
	Total_price float64
	Status      string
}

func (c Controller) CreateOrder(order Order, userPass string) (Order, error) {
	user, err := c.GetUser(order.User_id)
	if err != nil {
		return Order{}, fmt.Errorf("invalid user id")
	}
	if user.Password != userPass {
		return Order{}, fmt.Errorf("invalid user password")
	}

	queryStmt := `INSERT INTO orders (user_id, product_id, quantity, total_price, status)
                     VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	err = c.db.QueryRow(queryStmt,
		order.User_id,
		order.Product_id,
		order.Quantity,
		order.Total_price,
		order.Status).Scan(&order.Id)
	if err != nil {
		log.Println("failed to execute query:", err)
		return Order{}, err
	}

	return order, nil
}

func (c Controller) GetOrder(orderId string) (Order, error) {
	queryStmt := `SELECT id, user_id, product_id, quantity, total_price, status
                     FROM orders WHERE id = $1;`
	row := c.db.QueryRow(queryStmt, orderId)

	var order Order
	err := row.Scan(&order.Id, &order.User_id, &order.Product_id,
		&order.Quantity, &order.Total_price, &order.Status)
	if err != nil {
		return Order{}, err
	}

	return order, nil
}
