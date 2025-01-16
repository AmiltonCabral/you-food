package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	controlers "github.com/amiltoncabral/youFood/controllers"
)

type NewOrderReq struct {
	Order         controlers.Order
	User_password string
}

func (h Handler) OrderHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateOrder(w, r)
	case http.MethodGet:
		h.GetOrder(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var newOrderReq NewOrderReq
	err = json.Unmarshal(body, &newOrderReq)
	if err != nil {
		log.Println("failed to unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order, err := h.c.CreateOrder(newOrderReq.Order, newOrderReq.User_password)
	if err != nil {
		if err.Error() == "invalid user id" {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
			return
		}
		if err.Error() == "invalid product id" {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "product not found"})
			return
		}
		if err.Error() == "invalid user password" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println("failed to create order:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.sendOrderMessage(order)

	h.instrument.orderCounter.Add(r.Context(), 1)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func (h Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	order, err := h.c.GetOrder(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Println("failed to scan", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func (h Handler) sendOrderMessage(order controlers.Order) {
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(order)
	if err != nil {
		log.Println("failed to marshal order:", err)
		return
	}
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		log.Println("failed to publish a message:", err)
	}
}
