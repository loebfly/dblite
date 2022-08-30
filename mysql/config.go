package mysql

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var config = new(ymlConfig)

type ymlConfig struct {
	Mysql struct {
		Url   string `yml:"url"`
		Debug bool   `yml:"debug"`
		Pool  struct {
			Max     int `yml:"max"`
			Idle    int `yml:"idle"`
			Timeout struct {
				Idle int `yml:"idle"`
				Life int `yml:"life"`
			} `yml:"timeout"`
		}
	} `yml:"mysql"`
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
	if cfg.Mysql.Url == "" {
		return false
	}
	cfg.fillNull()
	return true
}
