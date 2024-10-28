package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type DeliveryMan struct {
    Id       string
    Name     string
    Password string
}

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

    var deliveryMan DeliveryMan
    err = json.Unmarshal(body, &deliveryMan)
    if err != nil {
        log.Println("failed to unmarshal:", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    queryStmt := `INSERT INTO delivery_man (id, name, password)
                  VALUES ($1, $2, $3) RETURNING id;`

    err = h.DB.QueryRow(queryStmt,
        deliveryMan.Id,
        deliveryMan.Name,
        deliveryMan.Password).Scan(&deliveryMan.Id)
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

    queryStmt := `SELECT id, name, password FROM delivery_man WHERE id = $1;`
    row := h.DB.QueryRow(queryStmt, id)

    var deliveryMan DeliveryMan
    err := row.Scan(&deliveryMan.Id, &deliveryMan.Name, &deliveryMan.Password)
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
