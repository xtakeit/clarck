package opt

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)
//TODO 基础库的东西太多了，暂时先不封装
type Base struct {
	//key string
	cli *redis.Client
}

func (b *Base) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	res := b.cli.Set(context.Background(), key, value, expiration)
	return res.Val(), res.Err()
}

func (b *Base) Del(keys ...string) (int64, error) {
	res := b.cli.Del(context.Background(), keys...)
	return res.Val(), res.Err()
}
