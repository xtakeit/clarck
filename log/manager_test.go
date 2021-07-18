package log

import (
	"testing"

	"github.com/anoxia/clarck/framework"
)

func TestInit(t *testing.T) {

	app := constructApp()
	err := Init(app)
	if err != nil {
		t.Error("日志组件初始化失败")
		return
	}
}

func constructApp() *framework.App {
	fileOutput := make(map[string]FileConfig)
	fileOutput[InfoLevel] = FileConfig{
		OutPath:     "./info",
		MaxAgeHours: 1,
		RotateHours: 1,
	}
	fileOutput[ErrorLevel] = FileConfig{
		OutPath:     "./error",
		MaxAgeHours: 1,
		RotateHours: 1,
	}

	app = &framework.App{}
	app.SetConfig(&Config{
		Log: LogConfigManager{
			Level:        InfoLevel,
			ReportCaller: false,
			LogType:      File,
			FileOutput:   fileOutput,
			RemoteOutput: nil,
		},
	})
	return app
}
