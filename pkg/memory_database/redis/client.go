package redis

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"github.com/t1pcrips/platform-pkg/pkg/memory_database"
	"time"
)

type rsClient struct {
	mastedDBC memory_database.DB
}

func NewClientRs(ctx context.Context, net string, dsn string, maxIdle int, maxIdleTimeout, ctxTimeout time.Duration) memory_database.Client {
	ctxWithTimeOut, _ := context.WithTimeout(ctx, ctxTimeout)

	//defer cancel()

	pool := &redis.Pool{
		IdleTimeout: maxIdleTimeout,
		MaxIdle:     maxIdle,
		DialContext: func(ctx context.Context) (redis.Conn, error) {
			return redis.DialContext(ctxWithTimeOut, net, dsn)
		},
	}

	return &rsClient{
		mastedDBC: NewRedis(pool),
	}
}

func (c *rsClient) DB() memory_database.DB {
	return c.mastedDBC
}

func (c *rsClient) Close() error {
	if c.mastedDBC != nil {
		c.mastedDBC.Close()
	}

	return nil
}
