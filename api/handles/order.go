package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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
