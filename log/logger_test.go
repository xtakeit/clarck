package log

import (
	"bufio"
	"encoding/json"
	"github.com/anoxia/clarck"
	"io"
	"os"
	"testing"
	"time"

	"github.com/anoxia/clarck/framework"
)

var app *framework.App

func TestMain(m *testing.M) {
	app = clarck.App()
	manager := New()
	app.Set("log", manager)

	m.Run()
}

func BenchmarkLog(b *testing.B) {
	output := make(map[string]interface{})
	output["data"] = "testInfo"

	for i := 0; i < b.N; i++ {
		Info(output)
	}

}

func BenchmarkLogParallel(b *testing.B) {
	output := make(map[string]interface{})
	output["data"] = "testInfo"

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Info(output)
		}
	})
}

func TestNewLogger(t *testing.T) {
	config := &Config{}
	app.ConfigLoad(config)

	_, err := NewLogger(config)
	if err != nil {
		t.Error("初始胡日志组件失败")
	}
}

func TestInfo(t *testing.T) {
	output := make(map[string]interface{})
	output["data"] = "testInfo"

	Info(output)

	validateLog(InfoLevel, InfoLevel, output, t)
}

func TestWarn(t *testing.T) {
	output := make(map[string]interface{})
	output["data"] = "testWarn"

	Warn(output)

	validateLog(ErrorLevel, WarnLevel, output, t)
}

func TestError(t *testing.T) {
	output := make(map[string]interface{})
	output["data"] = "testError"

	Error(output)

	validateLog(ErrorLevel, ErrorLevel, output, t)
}

// 验证输出到文件的日志内容,和实际内容是否一致
func validateLog(fileLevel string, level string, output map[string]interface{}, t *testing.T) {
	fileName := getFileName(fileLevel)
	line := readFinalLine(fileName, t)

	input := make(map[string]interface{})
	err := json.Unmarshal([]byte(line), &input)
	if err != nil {
		t.Error("unmarshal error")
		return
	}

	if output["data"] != input["data"] || input["level"] != level {
		t.Error("no match")
		return
	}
}

func getFileName(level string) string {
	return level + "." + time.Now().Format("2006-01-02-15")
}

func readFinalLine(filePath string, t *testing.T) string {
	_, err := os.Stat(filePath)
	if err != nil {
		t.Error("file not found")
	}

	file, err := os.Open(filePath)
	if err != nil {
		t.Error("clear file failed: ", err)
		return ""
	}
	defer file.Close()

	// 读取文件的最后一行
	var preLine, line string
	rd := bufio.NewReader(file)
	for {
		line, err = rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		preLine = line
	}
	return preLine
}
