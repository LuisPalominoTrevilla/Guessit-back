package redis

import "github.com/gomodule/redigo/redis"

// SetArbitraryPair inserts an arbitrary key value pair into redis
func (client *Client) SetArbitraryPair(key string, value interface{}) error {
	_, err := client.Conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}

// GetStringValue returns the string value associated with the provided key
func (client *Client) GetStringValue(key string) (interface{}, error) {
	res, err := redis.String(client.Conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return res, nil
}
