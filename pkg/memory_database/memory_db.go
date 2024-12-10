package memory_database

import (
	"context"
	"time"
)

type Client interface {
	DB() DB
	Close() error
}

type QueryExecute interface {
	DoContext(ctx context.Context, commandName string, timeout time.Duration, args ...interface{}) (interface{}, error)
}

type DB interface {
	QueryExecute
	Close()
	Ping(ctx context.Context, timeout time.Duration) error
}
