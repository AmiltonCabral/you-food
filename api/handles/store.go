package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	controlers "github.com/amiltoncabral/youFood/controllers"
)

func (h Handler) StoreHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateStore(w, r)
	case http.MethodGet:
		h.GetStore(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h Handler) CreateStore(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var store controlers.Store
	err = json.Unmarshal(body, &store)
	if err != nil {
		log.Println("failed to unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	store, err = h.c.CreateStore(store)
	if err != nil {
		log.Println("failed to execute query:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h Handler) GetStore(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	store, err := h.c.GetStore(id)

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
	json.NewEncoder(w).Encode(store)
}
