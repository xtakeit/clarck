package mysql

import (
	"github.com/anoxia/clarck"
	"gorm.io/gorm"
)

func init() {
	DbManager = &DatabaseManager{
		connections: make(map[string]connection),
	}
}

type DatabaseManager struct {
	connections map[string]connection
	appConfig   *clarck.Config
}

func (m *DatabaseManager) GetConnection(dbname ...string) (*gorm.DB, error) {
	if m.appConfig == nil {
		return nil, clarck.NewConfigError(-1, "appConfig.Database 为空")
	}
}

func (m *DatabaseManager) Exist(dbname string) bool {
	return true
}

func (m *DatabaseManager) SetConfig(config *clarck.Config) {
	DbManager.appConfig = config
}

// type MySQLConnections map[string]*gorm.DB

// var conns = make(MySQLConnections)
// var configs = NewConfigs()

// // 连接指定数据库（连接配置在 configs 中，见 config.go）
// func connect(dbname string) (db *gorm.DB, e error) {
// 	c, e := configs.Get(dbname)
// 	if e != nil {
// 		e = clarck.NewConfigError(-1, "数据库配置未找到")
// 		return
// 	}

// 	dsn := dsn(c)

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println("db connect failed: " + err.Error())
// 	}

// 	return
// }

// // 获取指定名称的数据库连接
// // 如果未传入名称则使用 default 作为默认目标数据库
// func Get(dbnames ...string) (db *gorm.DB, e error) {
// 	dbname := "default"
// 	for _, v := range dbnames {
// 		dbname = v
// 		break
// 	}

// 	if conns[dbname] == nil {
// 		db, e = connect(dbname)
// 		if e != nil {
// 			return
// 		} else {
// 			conns[dbname] = db
// 		}
// 	}

// 	db = conns[dbname]

// 	return
// }
