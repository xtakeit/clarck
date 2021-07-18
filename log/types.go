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

type LogConfigManager struct {
	Level        string                  // 日志级别
	ReportCaller bool                    // 是否打印调用者
	LogType      string                  // file、remote、all
	FileOutput   map[string]FileConfig   // 文件
	RemoteOutput map[string]RemoteConfig // 远程
}

const (
	File   = "file"
	Remote = "remote"
	All    = "all"
)
