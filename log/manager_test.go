package log

import (
	"github.com/anoxia/clarck/framework"
	"github.com/anoxia/clarck/types"
	"testing"
)

func TestInit(t *testing.T) {

	app := constructApp()
	err := Init(app)
	if err != nil {
		t.Error("日志组件初始化失败")
		return
	}
}

func constructApp () *framework.App {
	fileOutput := make(map[string]types.FileConfig)
	fileOutput[InfoLevel] = types.FileConfig{
		OutPath:     "./info",
		MaxAgeHours: 1,
		RotateHours: 1,
	}
	fileOutput[ErrorLevel] = types.FileConfig{
		OutPath:     "./error",
		MaxAgeHours: 1,
		RotateHours: 1,
	}

	app = &framework.App{}
	app.SetConfig(&types.Config{
		Log: types.LogConfigManager{
			Level:        InfoLevel,
			ReportCaller: false,
			LogType:      types.File,
			FileOutput:   fileOutput,
			RemoteOutput: nil,
		},
	})
	return app
}