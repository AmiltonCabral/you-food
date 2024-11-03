package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	controlers "github.com/amiltoncabral/youFood/controllers"
)

type NewProdReq struct {
	Product        controlers.Product
	Store_password string
}

func (h Handler) ProductHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateProduct(w, r)
	case http.MethodGet:
		h.GetProduct(w, r)
	case http.MethodPut:
		h.UpdateProduct(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var newProdReq NewProdReq
	err = json.Unmarshal(body, &newProdReq)
	if err != nil {
		log.Println("failed to unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.c.CreateProduct(newProdReq.Product, newProdReq.Store_password)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProdReq)
}

func (h Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	product, err := h.c.GetProduct(id)

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
	json.NewEncoder(w).Encode(product)
}

func (h Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "id is required"})
		return
	}

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var newProdReq NewProdReq
	err = json.Unmarshal(body, &newProdReq)
	if err != nil {
		log.Println("failed to unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.c.UpdateProduct(
		newProdReq.Product,
		newProdReq.Product.Store_id,
		newProdReq.Store_password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "product not found"})
			return
		}
		log.Println("failed to check product existence or execute update query:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
