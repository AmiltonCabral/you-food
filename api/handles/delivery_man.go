package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	controlers "github.com/amiltoncabral/youFood/controllers"
)

func (h Handler) DeliveryManHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateDeliveryMan(w, r)
	case http.MethodGet:
		h.GetDeliveryMan(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h Handler) CreateDeliveryMan(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var deliveryMan controlers.DeliveryMan
	err = json.Unmarshal(body, &deliveryMan)
	if err != nil {
		log.Println("failed to unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.c.GetDeliveryMan(deliveryMan.Id)
	if err != sql.ErrNoRows {
		w.WriteHeader(http.StatusConflict)
		return
	}

	deliveryMan, err = h.c.CreateDeliveryMan(deliveryMan)
	if err != nil {
		log.Println("failed to execute query:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(deliveryMan)
}

func (h Handler) GetDeliveryMan(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	deliveryMan, err := h.c.GetDeliveryMan(id)

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
	json.NewEncoder(w).Encode(deliveryMan)
}
