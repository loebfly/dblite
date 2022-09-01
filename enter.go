package dblite

import (
	"errors"
	"github.com/loebfly/dblite/mongo"
	"github.com/loebfly/dblite/mysql"
	"github.com/loebfly/dblite/redis"
)

var Mysql = new(mysql.Enter)
var Mongo = new(mongo.Enter)
var Redis = new(redis.Enter)

type Use string

const (
	UseMysql Use = "mysql"
	UseMongo Use = "mongo"
	UseRedis Use = "redis"
)

func Init(ymlPath string, use ...Use) error {
	if len(use) == 0 {
		return errors.New("use type is empty")
	}
	var errStr string
	for _, v := range use {
		switch v {
		case UseMysql:
			err := Mysql.Init(ymlPath)
			if err != nil {
				errStr += err.Error()
				errStr += ";"
			}
		case UseMongo:
			err := Mongo.Init(ymlPath)
			if err != nil {
				errStr += err.Error()
				errStr += ";"
			}
		case UseRedis:
			err := Redis.Init(ymlPath)
			if err != nil {
				errStr += err.Error()
				errStr += ";"
			}
		default:
			errStr += "use type is not support"
		}
	}
	if errStr != "" {
		return errors.New(errStr)
	}
	return nil
}

func SafeExit() {
	Mysql.SafeExit()
	Mongo.SafeExit()
	Redis.SafeExit()
}
