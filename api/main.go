package main

import (
	"github.com/amiltoncabral/youFood/database"
	"github.com/amiltoncabral/youFood/redis"
	"github.com/amiltoncabral/youFood/routes"
)

func main() {
	db := database.OpenConn()
	rc := redis.OpenConn()
	routes.HandleRequest(db)
	database.CloseConn(db)
	redis.CloseConn(rc)
}
