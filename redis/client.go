package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

)

type RedisClient struct {
	*redis.Client
}

func NewClient(cf ItemConfig) (cli *RedisClient, err error) {

	ocli := redis.NewClient(&redis.Options{
		Addr:     cf.GetAddr(),
		Password: cf.Password,
		DB:       cf.DB,
		PoolSize: cf.PoolSize,
	})

	ctx := context.Background()
	if err = ocli.Ping(ctx).Err(); err != nil {
		err = fmt.Errorf("client ping: %w", err)
		return
	}

	cli = &RedisClient{
		Client: ocli,
	}

	return
}

func (cli *RedisClient) Close() (err error) {
	if err = cli.Client.Close(); err != nil {
		err = fmt.Errorf("client close: %w", err)
		return
	}

	return
}
