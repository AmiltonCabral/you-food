package controllers

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Controller struct {
	Rmq_conn *amqp.Connection
	ctx      context.Context
}

func New(rmqConn *amqp.Connection) Controller {

	return Controller{rmqConn, context.Background()}
}
