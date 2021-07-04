package clarck

import "github.com/anoxia/clarck/db/mysql"

type Config struct {
	Database map[string]mysql.ConnectionConfig
}
