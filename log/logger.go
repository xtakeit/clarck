package log

import (
	"fmt"
	"io"
	"time"

	rotate "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"github.com/anoxia/clarck/types"
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

const (
	InfoLevel  = "info"
	WarnLevel  = "warning"
	ErrorLevel = "error"
	DebugLevel = "debug"
	FatalLevel = "fatal"
	PanicLevel = "panic"
)

// Logger 日志工具
type Logger struct {
	*logrus.Logger

	outMap map[string]io.Closer
}

// NewLogger 返回日志工具实例
func NewLogger(cf *types.LogConfigManager) (logger *Logger, err error) {
	logger = &Logger{
		Logger: logrus.New(),
		outMap: make(map[string]io.Closer),
	}

	// 设置标准输出格式
	logger.SetFormatter()

	// 关闭标准输出
	//logger.SetCloseStdOutput()

	// 设置是否开启调用者打印
	logger.SetReportCaller(cf)

	// 设置日志级别
	if err = logger.SetLevel(cf); err != nil {
		err = fmt.Errorf("set level: %w", err)
		return
	}

	// 设置勾子
	if err = logger.SetHook(cf); err != nil {
		err = fmt.Errorf("set hook: %w", err)
		return
	}

	return
}

func (logger *Logger) SetReportCaller(cf *types.LogConfigManager) {
	logger.Logger.SetReportCaller(cf.ReportCaller)
}

// SetFormatter 设置格式
func (logger *Logger) SetFormatter() {
	logger.Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: timeLayout,
	})
}

// SetLevel 设置级别
func (logger *Logger) SetLevel(cf *types.LogConfigManager) (err error) {
	level, err := logrus.ParseLevel(cf.Level)
	if err != nil {
		err = fmt.Errorf("parse level: %w", err)
		return
	}
	logger.Logger.SetLevel(level)
	return
}

/*func (logger *Logger) SetCloseStdOutput() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		err = fmt.Errorf("open os.DevNull: %w", err)
		return
	}

	writer := bufio.NewWriter(src)
	logger.SetOutput(writer)
}*/

// 设置日志输出，LogType不存在，默认文件输出
func (logger *Logger) SetHook(cf *types.LogConfigManager) (err error) {
	switch cf.LogType {
	case types.File:
		err = logger.setFileHook(cf.FileOutput)
	case types.Remote:
		//TODO 待添加远程输出
	case types.All:
		err = logger.setFileHook(cf.FileOutput)
		//TODO 待添加远程输出
	default:
		err = logger.setFileHook(cf.FileOutput)
	}
	return
}

func (logger *Logger) setFileHook(configs map[string]types.FileConfig) (err error) {

	// 创建日志输出回调，并防止被回收
	outMap := make(map[string]*rotate.RotateLogs)
	for k, v := range configs {
		out, err := rotate.New(
			v.OutPath+".%Y-%m-%d-%H", rotate.WithLinkName(v.OutPath),
			rotate.WithMaxAge(time.Duration(v.MaxAgeHours)*time.Hour),
			rotate.WithRotationTime(time.Duration(v.RotateHours)*time.Hour),
		)
		if err != nil {
			return fmt.Errorf("new %s rotate: %w", k, err)
		}
		logger.outMap[k] = out
		outMap[k] = out
	}

	// 根据不同的日志级别，调用设置不同的通道
	hook := lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: outMap[InfoLevel],
			logrus.InfoLevel:  outMap[InfoLevel],
			logrus.WarnLevel:  outMap[ErrorLevel],
			logrus.ErrorLevel: outMap[ErrorLevel],
			logrus.FatalLevel: outMap[ErrorLevel],
			logrus.PanicLevel: outMap[ErrorLevel],
		},
		&logrus.JSONFormatter{
			TimestampFormat: timeLayout,
		},
	)

	logger.Logger.AddHook(hook)
	return
}
