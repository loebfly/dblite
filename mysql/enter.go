package mysql

import "gorm.io/gorm"

type Enter struct{}

/*
LocalInit 本地初始化

示例

mysql:
  url: 连接地址(必配)
  pool: 可选
    max: 20 可选
    idle: 10 可选
    timeout: 可选
      idle: 60 可选
      life: 60 可选
*/
func (*Enter) LocalInit(ymlPath string) error {
	err := config.LocalInit(ymlPath)
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
func (*Enter) GetDB() (*gorm.DB, error) {
	return ctl.getDB()
}

// SafeExit 安全退出
func (*Enter) SafeExit() {
	ctl.disconnect()
}
