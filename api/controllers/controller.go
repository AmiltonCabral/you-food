package controllers

import (
	"context"
	"database/sql"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Controller struct {
	db              *sql.DB
	rd              *redis.Client
	ctx             context.Context
	cacheRefleshSec int
}

func New(db *sql.DB, rd *redis.Client) Controller {
	StrCacheRefleshSec := os.Getenv("CACHE_REFLESH_SEC")
	if StrCacheRefleshSec == "" {
		StrCacheRefleshSec = "30"
	}
	cacheRefleshSec, _ := strconv.Atoi(StrCacheRefleshSec)
	if cacheRefleshSec <= 0 {
		cacheRefleshSec = 30
	}

	return Controller{db, rd, context.Background(), cacheRefleshSec}
}
