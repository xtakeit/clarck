package clarck

import "github.com/anoxia/clarck/framework"

var app *framework.App

func App() *framework.App {
	if app == nil {
		app = framework.New()
	}
	return app
}

func Version() string {
	return "0.0.0"
}
