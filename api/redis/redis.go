package redis

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func OpenConn() *redis.Client {
	var err error

	password := "" // default, no password set
	if envPass := os.Getenv("REDIS_PASSWORD"); envPass != "" {
		password = envPass
	}

	dbNum := 0 // default DB
	if db := os.Getenv("REDIS_DB"); db != "" {
		dbNum, err = strconv.Atoi(db)
		if err != nil {
			log.Fatal(err)
		}
	}

	protocolNum := 2 // default connection protocol
	if protocol := os.Getenv("REDIS_PROTOCOL"); protocol != "" {
		protocolNum, err = strconv.Atoi(protocol)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Connecting to Redis at %s with db=%d protocol=%d\n",
		os.Getenv("REDIS_ADDR"), dbNum, protocolNum)

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: password,
		DB:       dbNum,
		Protocol: protocolNum,
	})

	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect redis: %v", err)
	}

	fmt.Printf("Successfully connected to redis! %s\n", pong)

	return client
}

func CloseConn(redisClient *redis.Client) {
	defer redisClient.Close()
}
