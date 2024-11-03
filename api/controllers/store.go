package controllers

type Store struct {
	Id       string
	Name     string
	Password string
	Address  string
}

func (c Controller) CreateStore(store Store) (Store, error) {
	queryStmt := `INSERT INTO stores (id, name, password, address)
	              VALUES ($1, $2, $3, $4);`

	err := c.db.QueryRow(queryStmt, store.Id, store.Name, store.Password, store.Address).Scan(&store.Id)

	if err != nil {
		return Store{}, err
	}

	return store, nil
}

func (c Controller) GetStore(storeId string) (Store, error) {
	queryStmt := `SELECT id, name, password, address FROM stores WHERE id = $1;`
	row := c.db.QueryRow(queryStmt, storeId)

	var store Store
	err := row.Scan(&store.Id, &store.Name, &store.Password, &store.Address)
	if err != nil {
		return Store{}, err
	}

	return store, nil
}
