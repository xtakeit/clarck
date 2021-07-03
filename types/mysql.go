package types

// MySQL 数据库连接配置
type MySQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Charset  string
}
