package controllers

import (
	"context"
	"database/sql"
	"os"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

type Controller struct {
	db              *sql.DB
	rd              *redis.Client
	Rmq_conn        *amqp.Connection
	ctx             context.Context
	cacheRefleshSec int
}

func New(db *sql.DB, rd *redis.Client, rmqConn *amqp.Connection) Controller {
	StrCacheRefleshSec := os.Getenv("CACHE_REFLESH_SEC")
	if StrCacheRefleshSec == "" {
		StrCacheRefleshSec = "30"
	}
	cacheRefleshSec, _ := strconv.Atoi(StrCacheRefleshSec)
	if cacheRefleshSec <= 0 {
		cacheRefleshSec = 30
	}

	return Controller{db, rd, rmqConn, context.Background(), cacheRefleshSec}
}
