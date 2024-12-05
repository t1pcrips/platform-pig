package redis

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"github.com/t1pcrips/platform-pkg/pkg/memory_database"
)

type rs struct {
	pool *redis.Pool
}

func NewRedis(pool *redis.Pool) memory_database.DB {
	return &rs{
		pool: pool,
	}
}

func (r *rs) DoContext(ctx context.Context, commandName string, args ...interface{}) (interface{}, error) {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect with redis")
	}

	defer func() {
		_ = r.pool.Close()
	}()

	reply, err := conn.Do(commandName, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Do in redis")
	}

	return reply, nil
}

func (r *rs) Ping(ctx context.Context) error {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = conn.Close()
	}()

	_, err = conn.Do("PING")
	if err != nil {
		return err
	}

	return nil
}

func (r *rs) Close() {
	r.pool.Close()
}
