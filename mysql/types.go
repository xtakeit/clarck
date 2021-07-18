package mysql

type ItemConfig struct {
	Name     string // 连接名称
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	Charset  string
}

type Config struct {
	Mysql map[string]ItemConfig
}
