package redis

import (
	"github.com/anoxia/clarck"
	"github.com/anoxia/clarck/errors"
	"github.com/anoxia/clarck/framework"
)

type Manager struct {
	clients map[string]*RedisClient
	app     *framework.App
	config  *Config
}

func New() *Manager {
	app := clarck.App()

	// 加载 MySQL 配置
	config := &Config{}
	app.ConfigLoad(config)

	return &Manager{
		clients: make(map[string]*RedisClient),
		app:     app,
		config:  config,
	}
}


func (m *Manager) Exist(name string) bool {
	if _, ok := m.clients[name]; ok {
		return true
	} else {
		return false
	}
}

func (m *Manager) GetClients(names ...string) (*RedisClient, error) {
	cli, err := m.GetConnection(names...)
	if err != nil {
		return nil, err
	}

	return cli, err
}

func (m *Manager) GetConnection(names ...string) (cli *RedisClient, err error) {
	name := nameWithDefault(names...)

	if m.Exist(name) {
		return m.clients[name], nil
	}

	if _, ok := m.config.Redis[name]; !ok {
		return nil, errors.NewConfigError(-1, "Redis数据库配置未找到，请检查配置文件(application.yml)", name)
	}


	config := m.config.Redis[name]

	// 创建数据库连接并保存
	// 下次可直接使用
	cli, err = NewClient(config)
	if err != nil {
		return nil, errors.NewFrameworkError(-1, "创建Redis数据库连接失败", err.Error())
	} else {
		m.clients[name] = cli
	}

	return
}

func nameWithDefault(names ...string) string {
	name := "default"
	for _, v := range names {
		name = v
		break
	}
	return name
}



