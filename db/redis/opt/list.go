package opt

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type List struct {
	key string
	cli *redis.Client
}

func (l *List) LIndex(index int64) (string, error) {
	res := l.cli.LIndex(context.Background(), l.key, index)
	return res.Val(), res.Err()
}

// LINSERT key BEFORE|AFTER pivot value
//
// 插入value 在 pivot 之前还是之后 (opt = BEFORE|AFTER)
func (l *List) LInsert(opt string, pivot, value interface{}) (int64, error) {
	res := l.cli.LInsert(context.Background(), l.key, opt, pivot, value)
	return res.Val(), res.Err()
}

func (l *List) LInsertBefore(pivot, value interface{}) (int64, error) {
	res := l.cli.LInsertBefore(context.Background(), l.key, pivot, value)
	return res.Val(), res.Err()
}

func (l *List) LInsertAfter(pivot, value interface{}) (int64, error) {
	res := l.cli.LInsertAfter(context.Background(), l.key, pivot, value)
	return res.Val(), res.Err()
}

func (l *List) LLen() (int64, error) {
	res := l.cli.LLen(context.Background(), l.key)
	return res.Val(), res.Err()
}

func (l *List) LPop() (string, error) {
	res := l.cli.LPop(context.Background(), l.key)
	return res.Val(), res.Err()
}

func (l *List) LPopCount(count int) ([]string, error) {
	res := l.cli.LPopCount(context.Background(), l.key, count)
	return res.Val(), res.Err()
}

func (l *List) LPos(value string, a redis.LPosArgs) (int64, error) {
	res := l.cli.LPos(context.Background(), l.key, value, a)
	return res.Val(), res.Err()
}

/*BLPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd
BRPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd
BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *StringCmd*/
