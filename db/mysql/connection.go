package mysql

import (
	"github.com/anoxia/clarck"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
	// 当前数据连接(nowConn)所使用的配置信息
	config ConnectionConfig

	// 用于存储旧配置对应的数据库连接
	// 起到平滑切换数据库连接的作用
	// 防止连接被释放时仍有业务在使用中
	old *gorm.DB
	// 用于存储新配置对应的数据库连接
	new *gorm.DB
}

func (c *Connection) Get() (*gorm.DB, error) {
	return c.new, nil
}

func (c *Connection) Create(config ConnectionConfig) (e error) {
	dsn := dsn(&config)

	c.new, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if e != nil {
		return clarck.NewFrameworkError(-1, "数据库连接失败", dsn, e.Error())
	}

	c.config = config

	return
}

func (c *Connection) Update(config ConnectionConfig) {
	c.old = c.new
	c.Create(config)
}

func (c *Connection) Free() {
	c.old = nil
	c.new = nil
}

func (c *Connection) ConfigIsDiff(config ConnectionConfig) bool {
	if c.config == config {
		return true
	}
	return false
}