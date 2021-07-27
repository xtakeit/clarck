package redis

import "strconv"

type ItemConfig struct {
	Host     string // 主机
	Password string // 密码
	DB       int    // 数据库
	Port     int    // 端口号
	PoolSize int    // 最大连接数
}

type Config struct {
	Redis map[string]ItemConfig
}

func (cf *ItemConfig) GetAddr() string {
	return cf.Host + ":" + strconv.FormatInt(int64(cf.Port), 10)
}
