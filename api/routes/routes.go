package routes

import (
	"database/sql"
	"log"
	"net/http"

	controllers "github.com/amiltoncabral/youFood/controllers"
	handlers "github.com/amiltoncabral/youFood/handles"
)

func HandleRequest(db *sql.DB) {
	c := controllers.New(db)
	h := handlers.New(c)
	http.HandleFunc("/user", h.UserHandler)
	http.HandleFunc("/product", h.ProductHandler)
	http.HandleFunc("/delivery-man", h.DeliveryManHandler)
	http.HandleFunc("/order", h.OrderHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
