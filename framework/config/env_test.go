package config

import (
	"fmt"
	"testing"

)

type ItemConfig struct {
	Name     string // 连接名称
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	Charset  string
}

type Config struct {
	Mysql map[string]ItemConfig
}

func TestMap(t *testing.T) {
	config := &Config{}

	LoadConfigFromFileEnv(".env", config)
	fmt.Println(config)
}

type FileConfig struct {
	OutPath     string // 输出路径
	MaxAgeHours int    // 错误日志文件最大保存小时数
	RotateHours int    // 错误日志文件轮转小时数
	Targets map[string]string
}

type RemoteConfig struct {
	Ip      string // ip地址
	Port    string // 端口号
	Api     string // 接口地址

}

type LogConfigManager struct {
	Log struct {
		Level        string                  // 日志级别
		ReportCaller bool                    // 是否打印调用者
		LogType      string                  // file、remote、all
		FileOutput   map[string]FileConfig   // 文件
		RemoteOutput map[string]RemoteConfig // 远程
	}
}

func TestComplex(t *testing.T) {
	config := &LogConfigManager{}

	LoadConfigFromFileEnv(".env", config)
	fmt.Println(config)
}

func TestToLowerFirst(t *testing.T) {
	name := "NameIsHaiXinX"
	name = toLowerFirst(name)
	fmt.Println(name)
}

func TestCamelToBaseLine(t *testing.T) {
	name := "NameIsHaiXinX"  //name_is_hai_xin_x
	name = camelToBaseLine(name)
	fmt.Println(name)
}

func TestCamelToLine(t *testing.T) {
	name := "MaxAgeHours"  //name_is_hai_xin_x
	name = camelToLine(name)
	fmt.Println(name)
}

