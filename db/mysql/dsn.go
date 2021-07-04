package mysql

import (
	"fmt"
)

// 生成 gorm 连接 DSN 字符串
func dsn(c *ConnectionConfig) string {
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
