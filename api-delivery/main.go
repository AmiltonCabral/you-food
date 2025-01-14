package main

import (
	"fmt"
	"os"
	"time"
	"youfood-delivery/routes"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	time.Sleep(5 * time.Second)
	rmq_conn, err := amqp.Dial(os.Getenv("RMQ_URL"))
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Connected to RabbitMQ\n")
	}
	defer rmq_conn.Close()

	routes.HandleRequest(rmq_conn)
}
