package dao

import (
	"context"
	"foragerServer/options"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	pool *redis.Pool
}

func NewRedis(c *options.RedisConfig) *Redis {
	return &Redis{
		pool: newPool(c),
	}
}

func newPool(c *options.RedisConfig) *redis.Pool {
	if c.Dial.DialTimeout <= 0 || c.Dial.ReadTimeout <= 0 || c.Dial.WriteTimeout <= 0 {
		panic("must config redis timeout")
	}
	ops := []redis.DialOption{
		redis.DialConnectTimeout(time.Duration(c.Dial.DialTimeout) * time.Millisecond),
		redis.DialReadTimeout(time.Duration(c.Dial.ReadTimeout) * time.Millisecond),
		redis.DialWriteTimeout(time.Duration(c.Dial.WriteTimeout) * time.Millisecond),
		redis.DialPassword(c.Dial.Password),
		redis.DialDatabase(c.Dial.Db),
	}
	pool := &redis.Pool{
		MaxIdle:   c.Pool.Idle,
		MaxActive: c.Pool.Active,
		// IdleTimeout: c.Pool.IdleTimeout,
		Wait: c.Pool.Wait,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(c.Dial.Protocol, c.Dial.Addr, ops...)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
	return pool
}

func (r *Redis) Do(ctx context.Context, commandName string, args ...interface{}) (reply interface{}, err error) {
	conn, err := r.pool.GetContext(ctx)
	defer conn.Close()
	reply, err = conn.Do(commandName, args...)
	return
}

func (r *Redis) Close() error {
	return r.pool.Close()
}

func (r *Redis) Conn(ctx context.Context) redis.Conn {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return nil
	}
	return conn
}
