package mysql

import (
	"github.com/anoxia/clarck"
	"github.com/anoxia/clarck/types"
)

type Configs struct {
	list map[string]*types.MySQLConfig
}

var configs *Configs

func (cs *Configs) Get(name string) (c *types.MySQLConfig, e error) {
	if cs.Exist(name) {
		c = cs.list[name]
		return
	}
	e = clarck.NewConfigError(1, "数据库配置不存在", name)
	return
}

func (cs *Configs) Exist(name string) bool {
	if _, ok := cs.list[name]; ok {
		return true
	}
	return false
}

func NewConfigs() *Configs {
	return &Configs{
		list: make(map[string]*types.MySQLConfig),
	}
}

// 新增数据库配置
func AddConfig(name string, c types.MySQLConfig) {
	if configs.Exist(name) {
		// 配置覆盖，记录日志
	}
	configs.list[name] = &c
}
