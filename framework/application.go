package framework

import (
	"io/ioutil"
	"os"

	"github.com/anoxia/clarck/types"
	"gopkg.in/yaml.v3"
)

type App struct {
	config    *types.Config
	event     map[string]Listenner
	bootstrap func(*App)
}

func New() (app *App) {
	app = &App{
		event: make(map[string]Listenner),
	}

	app.loadConfigFromFile(os.Getenv("APPLICATION_FILE"))

	return
}

func (app *App) Version() string {
	return "0.0.0"
}

func (app *App) Config() *types.Config {
	return app.config
}

func (app *App) RegisterListenner(name string, callback func(app *App, action string)) {
	app.event[name] = append(app.event[name], callback)
}

func (app *App) SetBootstrap(callback func(*App)) {
	app.bootstrap = callback
}

func (app *App) Run() {
	if app.bootstrap == nil {
		panic("服务启动失败，bootstrap 回掉方法未设置")
	}

	app.bootstrap(app)
}

func (app *App) loadConfigFromFile(filepath string) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic("配置文件[" + filepath + "]未找到")
	}

	config := types.Config{}
	yaml.Unmarshal(content, &config)

	app.config = &config
}
