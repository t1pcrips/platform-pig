package memory_database

import "context"

type Client interface {
	DB() DB
	Close() error
}

type QueryExecute interface {
	DoContext(ctx context.Context, commandName string, args ...interface{}) (interface{}, error)
}

type DB interface {
	QueryExecute
	Close()
}
