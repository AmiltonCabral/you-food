package controllers

import (
	"log"
)

type DeliveryMan struct {
	Id       string
	Name     string
	Password string
}

func (c Controller) CreateDeliveryMan(deliveryMan DeliveryMan) (DeliveryMan, error) {
	queryStmt := `INSERT INTO delivery_man (id, name, password)
                     VALUES ($1, $2, $3) RETURNING id;`

	err := c.db.QueryRow(queryStmt,
		deliveryMan.Id,
		deliveryMan.Name,
		deliveryMan.Password).Scan(&deliveryMan.Id)
	if err != nil {
		log.Println("failed to execute query:", err)
		return DeliveryMan{}, err
	}

	return deliveryMan, nil
}

func (c Controller) GetDeliveryMan(deliveryManId string) (DeliveryMan, error) {
	queryStmt := `SELECT id, name, password FROM delivery_man WHERE id = $1;`
	row := c.db.QueryRow(queryStmt, deliveryManId)

	var deliveryMan DeliveryMan
	err := row.Scan(&deliveryMan.Id, &deliveryMan.Name, &deliveryMan.Password)
	if err != nil {
		return DeliveryMan{}, err
	}

	return deliveryMan, nil
}
