package mysql

import (
	"fmt"

	"github.com/anoxia/clarck"
	"github.com/anoxia/clarck/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConnections map[string]*gorm.DB

var cons = make(MySQLConnections)

// 获取指定名称的数据库连接
// 如果未传入名称则使用 default 作为默认目标数据库
func Get(dbnames ...string) (db *gorm.DB, e error) {
	dbname := "default"
	for _, v := range dbnames {
		dbname = v
		break
	}

	if cons[dbname] == nil {
		db, e = connect(dbname)
		if e != nil {
			return
		} else {
			cons[dbname] = db
		}
	}

	db = cons[dbname]

	return
}

// 连接指定数据库（连接配置在 configs 中，见 config.go）
func connect(dbname string) (db *gorm.DB, e error) {
	c, e := configs.Get(dbname)
	if e != nil {
		e = clarck.New(1, "数据库配置未找到")
		return
	}

	dsn := dsn(c)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db connect failed: " + err.Error())
	}

	return
}

// 生成 gorm 连接 DSN 字符串
func dsn(c *types.MySQLConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		c.Charset,
	)
}
