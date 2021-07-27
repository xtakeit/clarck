package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// 加载配置文件
func Load() {
	// 配置文件路径
	filePath := os.Getenv("APPLICATION_FILE")

	paths := strings.Split(filePath, ".")
	suffix := paths[len(paths)-1]

	if !isSupportExt(suffix, viper.SupportedExts) {
		panic("无法加载此类配置文件：" + suffix)
	}

	viper.SetConfigType(suffix)
	viper.SetConfigFile(filePath)

	err := viper.ReadInConfig()
	if err != nil {
		panic("配置文件[" + filePath + "]未找到")
	}

}

// 加载配置到对象
// 默认加载驼峰格式 MyName，对应 myName
// 也可以配置  mapstructure: "anotherName" 进行配置
func LoadConfig(template interface{}) {
	err := viper.Unmarshal(template)
	if err != nil {
		panic(fmt.Sprintf("解析配置到对象失败：%+v", template))
	}
}

func isSupportExt(suffix string, extSlice []string) bool {
	for _, ext := range extSlice {
		if ext == suffix {
			return true
		}
	}
	return false
}
