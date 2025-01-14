package routes

import (
	"database/sql"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"

	controllers "github.com/amiltoncabral/youFood/controllers"
	handlers "github.com/amiltoncabral/youFood/handles"
)

func HandleRequest(db *sql.DB, rd *redis.Client, rmq_conn *amqp.Connection) {
	c := controllers.New(db, rd, rmq_conn)
	h := handlers.New(c)
	http.HandleFunc("/user", h.UserHandler)
	http.HandleFunc("/product", h.ProductHandler)
	http.HandleFunc("/delivery-man", h.DeliveryManHandler)
	http.HandleFunc("/order", h.OrderHandler)
	http.HandleFunc("/store", h.StoreHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
