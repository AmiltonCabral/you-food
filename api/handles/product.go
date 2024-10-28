package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Product struct {
	Id          int
	Store_id    string
	Name        string
	Description string
	Price       float64
	Amount      int
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

	var product Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		log.Println("failed to unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	queryStmt := `INSERT INTO products (store_id, name, description, price, ammount)
                  VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	err = h.DB.QueryRow(queryStmt,
		product.Store_id,
		product.Name,
		product.Description,
		product.Price,
		product.Amount).Scan(&product.Id)
	if err != nil {
		log.Println("failed to execute query:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (h Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	queryStmt := `SELECT id, store_id, name, description, price, ammount
                  FROM products WHERE id = $1;`
	row := h.DB.QueryRow(queryStmt, id)

	var product Product
	err := row.Scan(&product.Id, &product.Store_id, &product.Name,
		&product.Description, &product.Price, &product.Amount)
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

	var product Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		log.Println("failed to unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Primeiro verifica se o produto existe
	var existingProduct Product
	checkStmt := `SELECT id FROM products WHERE id = $1;`
	err = h.DB.QueryRow(checkStmt, id).Scan(&existingProduct.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "product not found"})
			return
		}
		log.Println("failed to check product existence:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Atualiza o produto
	queryStmt := `
        UPDATE products
        SET store_id = $1,
            name = $2,
            description = $3,
            price = $4,
            ammount = $5
        WHERE id = $6
        RETURNING id, store_id, name, description, price, ammount;`

	err = h.DB.QueryRow(
		queryStmt,
		product.Store_id,
		product.Name,
		product.Description,
		product.Price,
		product.Amount,
		id,
	).Scan(
		&product.Id,
		&product.Store_id,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Amount,
	)

	if err != nil {
		log.Println("failed to execute update query:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
