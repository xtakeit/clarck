package log

import (
	"github.com/sirupsen/logrus"

	"github.com/anoxia/clarck/framework"
)

type LogManager struct {
	loggers *Logger
	app     *framework.App
}

var manager *LogManager

func Init(app *framework.App) error {
	if manager == nil {
		log, err := NewLogger(&app.Config().Log)
		if err != nil {
			return err
		}

		manager = &LogManager{
			loggers: log,
			app:     app,
		}
		//TODO 等待添加文件监听事件
	}
	return nil
}

func (m *LogManager) Log(level logrus.Level, fields logrus.Fields) {

	entry := m.loggers.WithFields(fields)

	switch level {
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
	manager.Log(logrus.DebugLevel, fields)
}

func Info(fields map[string]interface{}) {
	manager.Log(logrus.InfoLevel, fields)
}

func Warn(fields map[string]interface{}) {
	manager.Log(logrus.WarnLevel, fields)
}

func Error(fields map[string]interface{}) {
	manager.Log(logrus.ErrorLevel, fields)
}

func Fatal(fields map[string]interface{}) {
	manager.Log(logrus.FatalLevel, fields)
}

func Panic(fields map[string]interface{}) {
	manager.Log(logrus.PanicLevel, fields)
}
