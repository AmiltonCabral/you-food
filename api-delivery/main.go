package main

import (
	"fmt"
	"os"
	"time"
	"youfood-delivery/routes"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	var rmq_conn *amqp.Connection
	var err error
	for i := 0; i < 5; i++ {
		rmq_conn, err = amqp.Dial(os.Getenv("RMQ_URL"))
		if err == nil {
			break
		}
		time.Sleep(6 * time.Second)
	}
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Connected to RabbitMQ\n")
	}
	defer rmq_conn.Close()

	routes.HandleRequest(rmq_conn)
}
