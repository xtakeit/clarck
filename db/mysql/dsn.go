package mysql

import (
	"fmt"

	"github.com/anoxia/clarck/types"
)

// 生成 gorm 连接 DSN 字符串
func DSN(c *types.MyqlConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Dbname,
		c.Charset,
	)
}
