package redis

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var manager *Manager

func TestMain(m *testing.M) {
	manager = New()
	m.Run()
}

func TestManager_GetClients(t *testing.T) {
	// 测试env中存在的配置
	redis, err := manager.GetClients("default")
	require.NotNil(t, redis, "获取redis链接失败%+v", err)

	// 测试env不存在的配置
	redis, err = manager.GetClients("unknow")
	require.Nil(t, redis)
}

func TestManager_GetConnection(t *testing.T) {
	redis, err := manager.GetConnection()
	require.NotNil(t, redis, "获取redis链接失败%+v", err)
}
