package dblite

import (
	"errors"
)

type UseType string

const (
	UseTypeMysql UseType = "mysql"
)

func LocalInit(ymlPath string, use ...UseType) error {
	if len(use) == 0 {
		return errors.New("use type is empty")
	}
	var errStr string
	for _, v := range use {
		switch v {
		case UseTypeMysql:
			err := Mysql.LocalInit(ymlPath)
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
