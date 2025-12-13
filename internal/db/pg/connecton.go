package pg

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Connection struct {
	connectionString string
}

func NewConnection(connectionString string) *Connection {
	return &Connection{
		connectionString: connectionString,
	}
}

func (conn *Connection) Connect() (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, conn.connectionString)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
