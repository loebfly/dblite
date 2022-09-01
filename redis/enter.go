package redis

import (
	"github.com/go-redis/redis"
	"github.com/loebfly/dblite/yml"
)

type Enter struct{}

/*
Init 本地初始化
redis:
  host: ip
  port: 端口
  password: 密码
  database: 库编号
  timeout: 1000
  pool:
    min: 3
    max: 20
    idle: 10
    timeout: 300
*/
func (*Enter) Init(ymlPath string) error {
	err := config.Init(ymlPath)
	if err != nil {
		return err
	}
	err = ctl.connect()
	if err != nil {
		return err
	}
	ctl.autoRetry()
	return nil
}

func (*Enter) InitObj(obj yml.Redis) error {
	err := config.InitObj(obj)
	if err != nil {
		return err
	}
	err = ctl.connect()
	if err != nil {
		return err
	}
	ctl.autoRetry()
	return nil
}

// Reconnect 重连
func (*Enter) Reconnect() error {
	return ctl.retry()
}

// GetDB 获取数据库
func (*Enter) GetDB() (*redis.Client, error) {
	return ctl.getDB()
}

// SafeExit 安全退出
func (*Enter) SafeExit() {
	ctl.disconnect()
}
