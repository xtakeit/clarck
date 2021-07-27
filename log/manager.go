package log

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/anoxia/clarck"
	"github.com/anoxia/clarck/framework"
)

type Manager struct {
	loggers *Logger
	app     *framework.App
	config  *Config
}

func New() *Manager {
	app := clarck.App()

	config := &Config{}
	app.ConfigLoad(config)

	logger, err := NewLogger(config)
	if err != nil {
		fmt.Println("logger加载失败")
	}

	manager := &Manager{
		loggers: logger,
		app:     app,
		config:  config,
	}

	return manager
}

func (m *Manager) Log(level logrus.Level, fields logrus.Fields) {

	entry := m.loggers.WithFields(fields)

	switch level {
	case logrus.DebugLevel:
		entry.Debug()
	case logrus.InfoLevel:
		entry.Info()
	case logrus.WarnLevel:
		entry.Warn()
	case logrus.ErrorLevel:
		entry.Error()
	case logrus.FatalLevel:
		entry.Fatal()
	case logrus.PanicLevel:
		entry.Panic()
	}
}

func Debug(fields map[string]interface{}) {
	manager := clarck.App().Get(componentName).(*Manager)
	manager.Log(logrus.DebugLevel, fields)
}

func Info(fields map[string]interface{}) {
	manager := clarck.App().Get(componentName).(*Manager)
	manager.Log(logrus.InfoLevel, fields)
}

func Warn(fields map[string]interface{}) {
	manager := clarck.App().Get(componentName).(*Manager)
	manager.Log(logrus.WarnLevel, fields)
}

func Error(fields map[string]interface{}) {
	manager := clarck.App().Get(componentName).(*Manager)
	manager.Log(logrus.ErrorLevel, fields)
}

func Fatal(fields map[string]interface{}) {
	manager := clarck.App().Get(componentName).(*Manager)
	manager.Log(logrus.FatalLevel, fields)
}

func Panic(fields map[string]interface{}) {
	manager := clarck.App().Get(componentName).(*Manager)
	manager.Log(logrus.PanicLevel, fields)
}
