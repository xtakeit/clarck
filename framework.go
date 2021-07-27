package clarck

import (
	"github.com/anoxia/clarck/framework"
	"github.com/anoxia/clarck/framework/config"
)

var app *framework.App

func App() *framework.App {
	if app == nil {
		config.Load()
		app = framework.New()
	}
	return app
}

func Version() string {
	return "0.0.0"
}
