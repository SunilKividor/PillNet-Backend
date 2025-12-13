package redisdb

import (
	"github.com/redis/go-redis/v9"
)

type Connection struct {
	connectionString string
}

func NewConnection(connectionString string) *Connection {
	return &Connection{
		connectionString: connectionString,
	}
}

func (conn *Connection) Connect() (*redis.Client, error) {
	opt, err := redis.ParseURL(conn.connectionString)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)
	return client, nil
}
