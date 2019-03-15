package redis

import (
	"github.com/gomodule/redigo/redis"
)

// SetArbitraryPair inserts an arbitrary key value pair into redis
func (client *Client) SetArbitraryPair(key string, value interface{}) error {
	_, err := client.Conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}

// SetExpArbitraryPair inserts an arbitrary key value pair into redis with expiration time
func (client *Client) SetExpArbitraryPair(key string, seconds int64, value interface{}) error {
	_, err := client.Conn.Do("SETEX", key, seconds, value)
	if err != nil {
		return err
	}
	return nil
}

// ExistsKey checks existance of key in redis
func (client *Client) ExistsKey(key string) (bool, error) {
	res, err := redis.Int(client.Conn.Do("EXISTS", key))
	if err != nil {
		return false, err
	}
	var exists bool
	if res == 1 {
		exists = true
	} else {
		exists = false
	}
	return exists, nil
}

// GetStringValue returns the string value associated with the provided key
func (client *Client) GetStringValue(key string) (interface{}, error) {
	res, err := redis.String(client.Conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return res, nil
}
