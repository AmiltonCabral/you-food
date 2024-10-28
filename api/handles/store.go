package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Store struct {
	Id       string
	Name     string
	Password string
	Address  string
}

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

	var store Store
	err = json.Unmarshal(body, &store)
	if err != nil {
		log.Println("failed to unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	queryStmt := `INSERT INTO stores (id, name, password, address)
	              VALUES ($1, $2, $3, $4);`

	err = h.DB.QueryRow(queryStmt, store.Id, store.Name, store.Password, store.Address).Scan(&store.Id)
	if err != nil {
		log.Println("failed to execute query:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h Handler) GetStore(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	queryStmt := `SELECT id, name, password, address FROM stores WHERE id = $1;`
	row := h.DB.QueryRow(queryStmt, id)

	var store Store
	err := row.Scan(&store.Id, &store.Name, &store.Password, &store.Address)
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
