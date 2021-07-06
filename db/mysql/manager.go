package mysql

import (
	"github.com/anoxia/clarck/errors"
	"github.com/anoxia/clarck/framework"
	"gorm.io/gorm"
)

type DatabaseManager struct {
	connections map[string]*Connection
	app         *framework.App
}

var manager *DatabaseManager

func Init(app *framework.App) {
	if manager == nil {
		manager = &DatabaseManager{
			connections: make(map[string]*Connection),
			app:         app,
		}

		app.RegisterListenner("update:config", func(app *framework.App, action string) {
		})
	}
}

func Manager() *DatabaseManager {
	return manager
}

func (m *DatabaseManager) Exist(name string) bool {
	if _, ok := m.connections[name]; ok {
		return true
	} else {
		return false
	}
}

func (m *DatabaseManager) GetDB(names ...string) (*gorm.DB, error) {
	connection, err := m.GetConnection(names...)
	if err != nil {
		return nil, err
	}

	return connection.GetDB(), err
}

func (m *DatabaseManager) GetConnection(names ...string) (conn *Connection, err error) {
	name := nameWithDefault(names...)

	if m.Exist(name) {
		return m.connections[name], nil
	}

	if _, ok := m.app.Config().Database.Mysql[name]; !ok {
		return nil, errors.NewConfigError(-1, "数据库配置未找到，请检查配置文件(application.yml)", name)
	}

	// 数据库连接配置
	// 来自 application.yml 配置文件
	config := m.app.Config().Database.Mysql[name]

	// 创建数据库连接并保存
	// 下次可直接使用
	conn, err = NewConnection(config)
	if err != nil {
		return nil, errors.NewFrameworkError(-1, "创建数据库连接失败", err.Error())
	} else {
		m.connections[name] = conn
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
