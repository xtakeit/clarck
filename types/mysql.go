package types

type MyqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	Charset  string
}

type DatabaseConfig struct {
	Mysql map[string]MyqlConfig
}
