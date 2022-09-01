package mongo

import (
	"github.com/loebfly/dblite/yml"
	"gopkg.in/mgo.v2"
)

type Enter struct{}

/*
Init 本地初始化
mongodb:
  url: "连接地址"
  database: "数据库"
  pool_max: 20
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

func (*Enter) InitObj(obj yml.Mongo) error {
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
func (*Enter) GetDB() (db *mgo.Database, closeDB func(), err error) {
	return ctl.getDB()
}

// SafeExit 安全退出
func (*Enter) SafeExit() {
	ctl.disconnect()
}
