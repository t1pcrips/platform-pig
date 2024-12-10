package redis

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"github.com/t1pcrips/platform-pkg/pkg/memory_database"
	"log"
	"time"
)

type handler func(ctx context.Context, conn redis.Conn) error

type rs struct {
	pool *redis.Pool
}

func NewRedis(pool *redis.Pool) memory_database.DB {
	return &rs{
		pool: pool,
	}
}

func (r *rs) DoContext(ctx context.Context, commandName string, timeout time.Duration, args ...interface{}) (interface{}, error) {
	var value interface{}
	err := r.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		var errEx error
		if commandName == "SET" {
			if len(args) >= 3 {
				ttl, ok := args[len(args)-1].(int)
				if ok {
					args = args[:len(args)-1]

					value, errEx = conn.Do("SETEX", append([]interface{}{args[0], ttl}, args[1:]...)...)
					if errEx != nil {
						return errEx
					}

					return nil
				}
			}
		}

		value, errEx = conn.Do(commandName, args...)
		if errEx != nil {
			return errors.Wrap(errEx, "failed to Do in redis")
		}

		return nil
	}, timeout)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (r *rs) Ping(ctx context.Context, timeout time.Duration) error {
	err := r.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		_, err := conn.Do("PING")
		if err != nil {
			return err
		}

		return nil
	}, timeout)

	if err != nil {
		return err
	}

	return nil
}

func (r *rs) Close() {
	r.pool.Close()
}

func (r *rs) execute(ctx context.Context, handler handler, timeout time.Duration) error {
	conn, err := r.getConnect(ctx, timeout)
	if err != nil {
		return err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Printf("failed to clise conn redis: %s\n", err.Error())
		}
	}()

	err = handler(ctx, conn)
	if err != nil {
		return err
	}

	return nil
}

func (r *rs) getConnect(ctx context.Context, timeout time.Duration) (redis.Conn, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := r.pool.GetContext(ctxWithTimeout)
	if err != nil {
		if conn != nil {
			_ = conn.Close()
		}
		return nil, err
	}

	return conn, nil
}
