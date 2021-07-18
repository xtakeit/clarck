package framework

import (
	"os"

	"github.com/anoxia/clarck/errors"
	"github.com/anoxia/clarck/framework/config"
)

type App struct {
	name        string
	module      string
	version     string
	configPath  string
	component   map[string]interface{}
	serverSetup func(*App)
}

func New() (app *App) {
	app = &App{
		component: make(map[string]interface{}),
	}

	// 配置文件路径
	app.configPath = os.Getenv("APPLICATION_FILE")

	// 加载配置
	c := &Config{}
	config.LoadConfigFromFile(app.configPath, c)

	// 将配置值注入当前实例
	app.version = c.Application.Version
	app.name = c.Application.Name
	app.module = c.Application.Module

	return
}

func (app *App) Name() string {
	return app.name
}

func (app *App) Module() string {
	return app.module
}

func (app *App) Version() string {
	return app.version
}

func (app *App) ConfigPath() string {
	return app.configPath
}

func (app *App) ConfigLoad(template interface{}) {
	config.LoadConfigFromFile(app.configPath, template)
}

func (app *App) Set(name string, module interface{}) {
	app.component[name] = module
}

func (app *App) Get(name string) interface{} {
	return app.component[name]
}

func (app *App) ServerSetup(callbacks ...func(*App)) {
	// 目前只支持单个 server 启动
	for _, callback := range callbacks {
		app.serverSetup = callback
		break
	}
}

func (app *App) Run() (err error) {
	if app.serverSetup != nil {
		app.serverSetup(app)
		return
	}

	err = errors.NewFrameworkError(-1, "服务启动失败，bootstrap 回掉方法未设置")
	return
}
