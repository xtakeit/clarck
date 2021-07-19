package redis

import "github.com/anoxia/clarck/framework"

type RedisManager struct {
	cli *Client
	app *framework.App
}

var manager *RedisManager

func Init(app *framework.App) error {
	if manager == nil {
		cli, err := NewClient(&app.Config().Redis)
		if err != nil {
			return err
		}
		manager = &RedisManager{
			cli: cli,
			app: app,
		}
		//TODO 等待添加监听配置文件事件
		//TODO 初始化外部的cache对象
	}
	return nil
}