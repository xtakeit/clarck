package mysql

// MySQL 数据库连接配置
type ConnectionConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Charset  string
}

// 数据库管理对象
var DbManager *DatabaseManager
