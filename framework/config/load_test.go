package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type ItemConfig struct {
	Name     string // 连接名称
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	Charset  string
}

type Config struct {
	Mysql map[string]ItemConfig
}

var testConfig = &Config{
	Mysql: map[string]ItemConfig{
		"default": {
			Name:     "test",
			Host:     "127.0.0.1",
			Port:     3306,
			User:     "root",
			Password: "123",
			Dbname:   "testDb",
			Charset:  "utf-8",
		},
	},
}

func TestMain(m *testing.M) {
	Load()
	m.Run()
}

func TestLoadConfig(t *testing.T) {
	config := &Config{}
	LoadConfig(config)
	require.Equal(t, testConfig, config)
}
