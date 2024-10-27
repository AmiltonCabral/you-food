package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/exp/rand"
)

type User struct {
	Id         string
	Name       string
	Password   string
	Order_code int
	Address    string
}

func (h Handler) UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateUser(w, r)
	case http.MethodGet:
		h.GetUser(w, r)
	case http.MethodPut:
		h.UpdateUser(w, r)
	case http.MethodDelete:
		h.DeleteUser(w, r)
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

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("failed to unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rand.Seed(uint64(time.Now().UnixNano()))
	user.Order_code = rand.Intn(9000) + 1000

	queryStmt := `INSERT INTO users (id, name, password, order_code, address)
                  VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	err = h.DB.QueryRow(queryStmt,
		user.Id,
		user.Name,
		user.Password,
		user.Order_code,
		user.Address).Scan(&user.Id)

	if err != nil {
		log.Println("failed to execute query:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	queryStmt := `SELECT id, name, password, order_code, address FROM users WHERE id = $1;`
	row := h.DB.QueryRow(queryStmt, id)

	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Password, &user.Order_code, &user.Address)
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

func (h Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {}

func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {}
