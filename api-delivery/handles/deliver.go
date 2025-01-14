package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Order struct {
	Id          int
	User_id     string
	Product_id  int
	Quantity    int
	Total_price float64
	Status      string
}

func (h Handler) DeliveryManHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetOrder(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	err := json.Unmarshal(h.getOrderMessage(), &order)
	if err != nil {
		log.Println("failed to unmarshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func (h Handler) getOrderMessage() []byte {
	ch, err := h.c.Rmq_conn.Channel()
	if err != nil {
		log.Println("failed to open a channel:", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"orders", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Println("failed to declare a queue:", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Println("failed to register a consumer:", err)
	}

	msg := <-msgs

	return msg.Body
}
