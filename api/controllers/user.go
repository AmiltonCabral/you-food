package controllers

import (
	"log"
	"time"

	"golang.org/x/exp/rand"
)

type User struct {
	Id         string
	Name       string
	Password   string
	Order_code int
	Address    string
}

func (c Controller) CreateUser(user User) (User, error) {
	rand.Seed(uint64(time.Now().UnixNano()))
	user.Order_code = rand.Intn(9000) + 1000
	queryStmt := `INSERT INTO users (id, name, password, order_code, address)
                     VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	_, err := c.db.Exec(queryStmt,
		user.Id,
		user.Name,
		user.Password,
		user.Order_code,
		user.Address)
	if err != nil {
		log.Println("failed to execute query:", err)
		return User{}, err
	}

	return user, nil
}

func (c Controller) GetUser(userId string) (User, error) {
	queryStmt := `SELECT id, name, password, order_code, address FROM users WHERE id = $1;`
	row := c.db.QueryRow(queryStmt, userId)

	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Password, &user.Order_code, &user.Address)

	if err != nil {
		return User{}, err
	}

	return user, nil
}
