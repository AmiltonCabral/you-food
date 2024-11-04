package controllers

import (
	"fmt"
	"log"
)

type Product struct {
	Id          int
	Store_id    string
	Name        string
	Description string
	Price       float64
	Amount      int
}

func (c Controller) CreateProduct(product Product, password string) (Product, error) {
	store, err := c.GetStore(product.Store_id)
	if store.Password != password {
		return Product{}, fmt.Errorf("invalid store password")
	}

	queryStmt := `INSERT INTO products (store_id, name, description, price, ammount)
                     VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	err = c.db.QueryRow(queryStmt,
		product.Store_id,
		product.Name,
		product.Description,
		product.Price,
		product.Amount).Scan(&product.Id)
	if err != nil {
		log.Println("failed to execute query:", err)
		return Product{}, err
	}

	return product, nil
}

func (c Controller) GetProduct(productId string) (Product, error) {
	queryStmt := `SELECT id, store_id, name, description, price, ammount
                     FROM products WHERE id = $1;`
	row := c.db.QueryRow(queryStmt, productId)

	var product Product
	err := row.Scan(&product.Id, &product.Store_id, &product.Name,
		&product.Description, &product.Price, &product.Amount)

	return product, err
}

func (c Controller) updateProduct(product Product) (Product, error) {
	// Primeiro verifica se o produto existe
	var existingProduct Product
	checkStmt := `SELECT id FROM products WHERE id = $1;`
	err := c.db.QueryRow(checkStmt, product.Id).Scan(&existingProduct.Id)
	if err != nil {
		return Product{}, err
	}

	queryStmt := `
           UPDATE products
           SET store_id = $1,
               name = $2,
               description = $3,
               price = $4,
               ammount = $5
           WHERE id = $6
           RETURNING id, store_id, name, description, price, ammount;`

	err = c.db.QueryRow(
		queryStmt,
		product.Store_id,
		product.Name,
		product.Description,
		product.Price,
		product.Amount,
		product.Id,
	).Scan(
		&product.Id,
		&product.Store_id,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Amount,
	)

	return product, err
}

func (c Controller) UpdateProduct(product Product, storeId string, storePassword string) (Product, error) {
	store, err := c.GetStore(storeId)
	if err != nil {
		return Product{}, fmt.Errorf("invalid store id")
	}
	if store.Password != storePassword {
		return Product{}, fmt.Errorf("invalid store password")
	}

	product, err = c.updateProduct(product)
	if err != nil {
		return Product{}, fmt.Errorf("failed to update product")
	}

	return product, nil
}

func (c Controller) BuyProduct(productId string, amount int) (Product, error) {
	var product Product
	err := c.db.QueryRow(`SELECT id, store_id, name, description, price, ammount
                     FROM products WHERE id = $1;`, productId).Scan(&product.Id, &product.Store_id, &product.Name,
		&product.Description, &product.Price, &product.Amount)
	if err != nil {
		return Product{}, err
	}

	if product.Amount < amount {
		return Product{}, fmt.Errorf("insufficient product amount")
	}

	product.Amount -= amount

	queryStmt := `UPDATE products SET ammount = $1 WHERE id = $2;`
	_, err = c.db.Exec(queryStmt, product.Amount, product.Id)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (c Controller) SearchProducts(query string) ([]Product, error) {
	queryStmt := `
        SELECT id, store_id, name, description, price, ammount
        FROM products
        WHERE name ILIKE $1 OR description ILIKE $1;
    `
	// O ILIKE faz uma busca case-insensitive e o % permite match parcial
	searchPattern := "%" + query + "%"

	rows, err := c.db.Query(queryStmt, searchPattern)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.Id,
			&product.Store_id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Amount,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
