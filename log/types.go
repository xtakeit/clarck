package log

type FileConfig struct {
	OutPath     string // 输出路径
	MaxAgeHours int    // 错误日志文件最大保存小时数
	RotateHours int    // 错误日志文件轮转小时数
}

type RemoteConfig struct {
	Ip   string // ip地址
	Port string // 端口号
	Api  string // 接口地址
}

type Config struct {
	Log struct {
		Level        string
		ReportCaller bool
		LogType      string
		FileOutput   map[string]FileConfig
		RemoteOutput map[string]RemoteConfig
	}
}

const (
	File   = "file"
	Remote = "remote"
	All    = "all"
)
