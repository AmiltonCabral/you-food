package main

import (
	"fmt"
	"os"

	"github.com/amiltoncabral/youFood/database"
	"github.com/amiltoncabral/youFood/redis"
	"github.com/amiltoncabral/youFood/routes"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db := database.OpenConn()
	defer database.CloseConn(db)

	rd := redis.OpenConn()
	defer redis.CloseConn(rd)

	rmq_conn, err := amqp.Dial(os.Getenv("RMQ_URL"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Connected to RabbitMQ\n")
	}
	defer rmq_conn.Close()

	routes.HandleRequest(db, rd, rmq_conn)
}
