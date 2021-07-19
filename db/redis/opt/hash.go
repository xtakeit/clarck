package opt

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Hash struct {
	key string
	cli *redis.Client
}

func (h *Hash) HGet(fields string) (string, error) {
	res := h.cli.HGet(context.Background(), h.key, fields)
	return res.Val(), res.Err()
}

func (h *Hash) HSet(values ...interface{}) (int64, error) {
	res := h.cli.HSet(context.Background(), h.key, values)
	return res.Val(), res.Err()
}

func (h *Hash) HDel(fields ...string) (int64, error) {
	res := h.cli.HDel(context.Background(), h.key, fields...)
	return res.Val(), res.Err()
}

func (h *Hash) HMGet(fields ...string) ([]interface{}, error) {
	res := h.cli.HMGet(context.Background(), h.key, fields...)
	return res.Val(), res.Err()
}

func (h *Hash) HExists(field string) (bool, error) {
	res := h.cli.HExists(context.Background(), h.key, field)
	return res.Val(), res.Err()
}

func (h *Hash) HGetAll() (map[string]string, error) {
	res := h.cli.HGetAll(context.Background(), h.key)
	return res.Val(), res.Err()
}

func (h *Hash) HIncrBy(field string, incr int64) (int64, error) {
	res := h.cli.HIncrBy(context.Background(), h.key, field, incr)
	return res.Val(), res.Err()
}

func (h *Hash) HIncrByFloat(field string, incr float64) (float64, error) {
	res := h.cli.HIncrByFloat(context.Background(), h.key, field, incr)
	return res.Val(), res.Err()
}

func (h *Hash) HKeys() ([]string, error) {
	res := h.cli.HKeys(context.Background(), h.key)
	return res.Val(), res.Err()
}

func (h *Hash) HLen() (int64, error) {
	res := h.cli.HLen(context.Background(), h.key)
	return res.Val(), res.Err()
}

// HRandField redis-server version >= 6.2.0.
func (h *Hash) HRandField(count int, withValues bool) ([]string, error) {
	res := h.cli.HRandField(context.Background(), h.key, count, withValues)
	return res.Val(), res.Err()
}

//返回顺序为： keys、redisCursor、error
func (h *Hash) HScan(cursor uint64, match string, count int64) (keys []string, redisCursor uint64, err error) {
	res := h.cli.HScan(context.Background(), h.key, cursor, match, count)
	res1, res2 := res.Val()
	return res1, res2, res.Err()
}

func (h *Hash) HSetNX(field string, value interface{}) (bool, error) {
	res := h.cli.HSetNX(context.Background(), h.key, field, value)
	return res.Val(), res.Err()
}

func (h *Hash) HVals() ([]string, error) {
	res := h.cli.HVals(context.Background(), h.key)
	return res.Val(), res.Err()
}

// 弃用的方法
func (h *Hash) hMSet(keyValuePair ...interface{}) {
}
