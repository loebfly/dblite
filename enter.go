package dblite

import (
	"dblite/mysql"
	"errors"
)

var Mysql = new(mysql.Enter)

type UseType string

const (
	UseTypeMysql UseType = "mysql"
)

func Init(ymlPath string, use ...UseType) error {
	if len(use) == 0 {
		return errors.New("use type is empty")
	}
	var errStr string
	for _, v := range use {
		switch v {
		case UseTypeMysql:
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
