package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
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

    var order Order
    err = json.Unmarshal(body, &order)
    if err != nil {
        log.Println("failed to unmarshal:", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    queryStmt := `INSERT INTO orders (user_id, product_id, quantity, total_price, status)
                  VALUES ($1, $2, $3, $4, $5) RETURNING id;`

    err = h.DB.QueryRow(queryStmt,
        order.User_id,
        order.Product_id,
        order.Quantity,
        order.Total_price,
        order.Status).Scan(&order.Id)
    if err != nil {
        log.Println("failed to execute query:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(order)
}

func (h Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")

    queryStmt := `SELECT id, user_id, product_id, quantity, total_price, status
                  FROM orders WHERE id = $1;`
    row := h.DB.QueryRow(queryStmt, id)

    var order Order
    err := row.Scan(&order.Id, &order.User_id, &order.Product_id,
                    &order.Quantity, &order.Total_price, &order.Status)
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
