package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func LoadConfigFromFile(filepath string, template interface{}) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic("配置文件[" + filepath + "]未找到")
	}

	yaml.Unmarshal(content, template)
}
