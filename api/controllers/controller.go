package controllers

import "database/sql"

type Controller struct {
	db *sql.DB
}

func New(db *sql.DB) Controller {
	return Controller{db}
}
