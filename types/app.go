package types

type Http struct {
	Addr string
}

type Config struct {
	Database DatabaseConfig
	Redis    RedisConf
	Http     Http
	Log      LogConfigManager
}
