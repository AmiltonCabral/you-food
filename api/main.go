package main

import (
	"github.com/amiltoncabral/youFood/database"
	"github.com/amiltoncabral/youFood/redis"
	"github.com/amiltoncabral/youFood/routes"
)

func main() {
	db := database.OpenConn()
	rd := redis.OpenConn()
	routes.HandleRequest(db, rd)
	database.CloseConn(db)
	redis.CloseConn(rd)
}
