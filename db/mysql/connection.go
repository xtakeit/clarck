package mysql

import (
	"github.com/anoxia/clarck/errors"
	"github.com/anoxia/clarck/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
	config types.MyqlConfig
	// old    *gorm.DB
	new *gorm.DB
}

func (c *Connection) GetDB() *gorm.DB {
	return c.new
}

func (c *Connection) open() error {
	dsn := DSN(&c.config)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.NewFrameworkError(-1, "数据库连接失败", err.Error())
	} else {
		c.new = db
	}

	return err
}

// func (c *Connection) update(config ConnectionConfig) {
// 	c.old = c.new
// 	c.config = config
// 	c.open()
// }

func (c *Connection) Free() {
	// c.old = nil
	c.new = nil
}

func NewConnection(config types.MyqlConfig) (*Connection, error) {
	con := Connection{
		config: config,
	}

	err := con.open()

	return &con, err
}
