package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	controlers "github.com/amiltoncabral/youFood/controllers"
)

func (h Handler) UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateUser(w, r)
	case http.MethodGet:
		h.GetUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user controlers.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("failed to unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err = h.c.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	user, err := h.c.GetUser(id)

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
	json.NewEncoder(w).Encode(user)
}
