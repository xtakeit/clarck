package types

type Http struct {
	Addr string
}

type Config struct {
	Database DatabaseConfig
	Http     Http
	Log      LogConfigManager
}
