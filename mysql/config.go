package mysql

import (
	"errors"
	"github.com/loebfly/dblite/yml"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var config = new(ymlConfig)

type ymlConfig struct {
	Mysql yml.Mysql `yaml:"mysql"`
}

func (cfg *ymlConfig) Init(ymlPath string) error {
	file, err := ioutil.ReadFile(ymlPath)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(file, cfg); err != nil {
		return err
	}
	if cfg.Mysql.Url == "" {
		return errors.New("mysql.url not null")
	}
	cfg.fillNull()
	return nil
}

func (cfg *ymlConfig) InitObj(obj yml.Mysql) error {
	if obj.Url == "" {
		return errors.New("mysql.url not null")
	}
	cfg.Mysql = obj
	cfg.fillNull()
	return nil
}

func (cfg *ymlConfig) fillNull() {
	if cfg.Mysql.Pool.Max == 0 {
		cfg.Mysql.Pool.Max = 20
	}
	if cfg.Mysql.Pool.Idle == 0 {
		cfg.Mysql.Pool.Idle = 10
	}
	if cfg.Mysql.Pool.Timeout.Life == 0 {
		cfg.Mysql.Pool.Timeout.Life = 60
	}
	if cfg.Mysql.Pool.Timeout.Idle == 0 {
		cfg.Mysql.Pool.Timeout.Idle = 60
	}
}

func (cfg *ymlConfig) isCorrect() bool {
	return cfg.Mysql.Url != ""
}
