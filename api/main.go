package main

import (
	"github.com/amiltoncabral/youFood/database"
	"github.com/amiltoncabral/youFood/routes"
)

func main() {
	db := database.OpenConn()
	routes.HandleRequest(db)
	database.CloseConn(db)
}
