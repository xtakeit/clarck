package framework

import (
	"os"

	"github.com/anoxia/clarck/errors"
	"github.com/anoxia/clarck/framework/config"
)

type App struct {
	configPath  string
	info        *Info
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
	info := &Info{}
	config.LoadConfig(info)
	app.info = info

	return
}

func (app *App) Name() string {
	return app.info.Application.Name
}

func (app *App) Module() string {
	return app.info.Application.Module
}

func (app *App) Version() string {
	return app.info.Application.Version
}

func (app *App) ConfigPath() string {
	return app.configPath
}

func (app *App) ConfigLoad(template interface{}) {
	config.LoadConfig(template)
}
//TODO 建议添加类型检查和值检擦，防止组件间key覆盖（或输出提示）
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
