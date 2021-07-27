package kafka

import (
	"github.com/anoxia/clarck"
	"github.com/anoxia/clarck/framework"
)

type Manager struct {
	producers map[string]*SyncProducer
	consumers map[string]*ConsumerGroup
	app       *framework.App
	config    *Config
}

func New() *Manager {
	app := clarck.App()

	config := &Config{}
	app.ConfigLoad(config)

	return &Manager{
		producers: make(map[string]*SyncProducer),
		consumers: make(map[string]*ConsumerGroup),
		app:       app,
		config:    config,
	}
}


