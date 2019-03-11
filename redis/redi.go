package redis

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

// Client serves as a wrapper for all operations
type Client struct {
	Conn redis.Conn
}

// InitRedis initializes the database, returns a client
func InitRedis() *Client {
	redisDB, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		log.Fatal(err)
	}

	return &Client{
		Conn: redisDB,
	}
}
