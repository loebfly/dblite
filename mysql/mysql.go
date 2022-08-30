package mysql

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var ctl = new(control)

type control struct {
	gormDB *gorm.DB
}

func (c *control) connect() error {
	gormDB, err := gorm.Open(mysql.Open(config.Mysql.Url), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxOpenConns(config.Mysql.Pool.Max)
	sqlDB.SetMaxIdleConns(config.Mysql.Pool.Idle)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.Mysql.Pool.Timeout.Idle) * time.Second)
	sqlDB.SetConnMaxLifetime(time.Duration(config.Mysql.Pool.Timeout.Life) * time.Minute)
	if config.Mysql.Debug {
		c.gormDB = gormDB.Debug()
	} else {
		c.gormDB = gormDB
	}
	return nil
}

func (c *control) disconnect() {
	if c.gormDB != nil {
		db, _ := c.gormDB.DB()
		_ = db.Close()
		c.gormDB = nil
	}
}

func (c *control) retry() error {
	if !config.isCorrect() {
		return errors.New("retry connect failed, config is not correct")
	}
	if c.gormDB == nil {
		err := c.connect()
		if err != nil {
			return errors.New("retry connect failed, " + err.Error())
		}
	}
	db, _ := c.gormDB.DB()
	err := db.Ping()
	if err != nil {
		c.disconnect()
		err = c.connect()
		if err != nil {
			return errors.New("retry connect failed, " + err.Error())
		}
	}
	return nil
}

func (c *control) autoRetry() {
	//设置定时任务自动检查
	ticker := time.NewTicker(time.Minute * 30)
	go func(c *control) {
		for range ticker.C {
			_ = c.retry()
		}
	}(c)
}

func (c *control) getDB() (*gorm.DB, error) {
	if c.gormDB == nil {
		err := c.connect()
		if err != nil {
			return nil, err
		}
	}
	return c.gormDB, nil
}
