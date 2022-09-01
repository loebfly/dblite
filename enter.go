package dblite

import (
	"errors"
	"github.com/loebfly/dblite/mysql"
)

var Mysql = new(mysql.Enter)

type Use string

const (
	UseMysql Use = "mysql"
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
		default:
			errStr += "use type is not support"
		}
	}
	if errStr != "" {
		return errors.New(errStr)
	}
	return nil
}
