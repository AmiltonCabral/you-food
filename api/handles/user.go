package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/exp/rand"
)

type User struct {
	id         string
	name       string
	password   string
	order_code int
	address    string
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
	// Read to request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(500)
		return
	}
	var user User
	json.Unmarshal(body, &user)

	// user.id = (uuid.New()).String()

	rand.Seed(uint64(time.Now().UnixNano()))
	user.order_code = rand.Intn(9000) + 1000

	queryStmt := `INSERT INTO users (id,name,password,order_code,address) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	err = h.DB.QueryRow(queryStmt, &user.id, &user.name, &user.password, &user.order_code, &user.address).Scan(&user.id)
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (h Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	queryStmt := `SELECT * FROM users WHERE id = $1 ;`
	results, err := h.DB.Query(queryStmt, id)
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	var user User
	for results.Next() {
		err = results.Scan(&user.id, &user.name, &user.password, &user.order_code, &user.address)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(500)
			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {}

func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {}
