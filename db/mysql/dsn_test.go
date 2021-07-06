package mysql

import "testing"

var configForTest = ConnectionConfig{
	Host:     "127.0.0.1",
	Port:     3306,
	User:     "root",
	Password: "123456",
	Database: "db_test",
	Charset:  "utf8",
}

func TestDsn(t *testing.T) {
	dsn := DSN(&configForTest)

	if dsn != "root:123456@tcp(127.0.0.1:3306)/db_test?charset=utf8" {
		t.Error("生成 dsn 返回字符串不符合预期")
	}
}
