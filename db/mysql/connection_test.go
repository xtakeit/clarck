package mysql

import "testing"

func TestNewConnection(t *testing.T) {
	_, err := NewConnection(configForTest)
	if err != nil {
		t.Error("创建数据库连接失败")
	}
}

func TestGetDB(t *testing.T) {
	con, _ := NewConnection(configForTest)
	_, err := con.GetDB()
	if err != nil {
		t.Error("获取数据库链接失败")
	}
}
