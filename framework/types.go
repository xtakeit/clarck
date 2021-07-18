package framework

type Config struct {
	Application struct {
		Name    string // 应用名称
		Module  string // 应用所属模块
		Version string // 应用版本
	}
}
