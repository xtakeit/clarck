package redis

import (
	"context"
	"fmt"
	"github.com/anoxia/clarck/types"
	"github.com/go-redis/redis/v8"
)

type Client struct {
	*redis.Client
}

func NewClient(cf *types.RedisConf) (cli *Client, err error) {

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

	cli = &Client{
		Client: ocli,
	}

	return
}

func (cli *Client) Close() (err error) {
	if err = cli.Client.Close(); err != nil {
		err = fmt.Errorf("client close: %w", err)
		return
	}

	return
}
