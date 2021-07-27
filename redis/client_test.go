package redis

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var clientConfigForTest = ItemConfig{
	Host:     "127.0.0.1",
	Password: "",
	DB:       0,
	Port:     6379,
	PoolSize: 10,
}

func TestRedisClient(t *testing.T) {

	cli, err := NewClient(clientConfigForTest)
	require.NotNil(t, cli, "创建redis连接错误：%+v", err)

	key, val := "test", "testVal"
	cmd := cli.Set(context.Background(), key, val, 5*time.Minute)
	require.Nil(t, cmd.Err(), "设置redis值失败: %+v", cmd.Err())

	resCmd := cli.Get(context.Background(), "test")
	require.Nil(t, resCmd.Err(), "设置redis值失败: %+v", resCmd.Err())
	require.Equal(t, val, resCmd.Val())

	err = cli.Close()
	require.Nil(t, err, "关闭数据库错误")
}
