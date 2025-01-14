package routes

import (
	"log"
	"net/http"
	"youfood-delivery/controllers"

	amqp "github.com/rabbitmq/amqp091-go"

	handlers "youfood-delivery/handles"
)

func HandleRequest(rmq_conn *amqp.Connection) {
	c := controllers.New(rmq_conn)
	h := handlers.New(c)
	http.HandleFunc("/delivery-man", h.DeliveryManHandler)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
