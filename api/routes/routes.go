package routes

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"

	controllers "github.com/amiltoncabral/youFood/controllers"
	handlers "github.com/amiltoncabral/youFood/handles"
)

func HandleRequest(db *sql.DB, rd *redis.Client) {
	c := controllers.New(db, rd)
	h := handlers.New(c)
	http.HandleFunc("/user", h.UserHandler)
	http.HandleFunc("/product", h.ProductHandler)
	http.HandleFunc("/delivery-man", h.DeliveryManHandler)
	http.HandleFunc("/order", h.OrderHandler)
	http.HandleFunc("/store", h.StoreHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
