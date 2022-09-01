package redis

import (
	"errors"
	"github.com/loebfly/dblite/yml"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var config = new(ymlConfig)

type ymlConfig struct {
	Redis yml.Redis `yaml:"redis"`
}

func (cfg *ymlConfig) Init(ymlPath string) error {
	file, err := ioutil.ReadFile(ymlPath)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(file, cfg); err != nil {
		return err
	}
	if err := cfg.checkObj(cfg.Redis); err != nil {
		return err
	}
	cfg.fillNull()
	return nil
}

func (cfg *ymlConfig) InitObj(obj yml.Redis) error {
	if err := cfg.checkObj(obj); err != nil {
		return err
	}
	cfg.Redis = obj
	cfg.fillNull()
	return nil
}

func (cfg *ymlConfig) checkObj(obj yml.Redis) error {
	if obj.Host == "" {
		return errors.New("redis.host not null")
	}
	if obj.Port == 0 {
		return errors.New("redis.port not null")
	}
	if obj.Password == "" {
		return errors.New("redis.password not null")
	}
	return nil
}

func (cfg *ymlConfig) fillNull() {
	if cfg.Redis.Timeout == 0 {
		cfg.Redis.Timeout = 1000
	}
	if cfg.Redis.Pool.Min == 0 {
		cfg.Redis.Pool.Min = 3
	}
	if cfg.Redis.Pool.Max == 0 {
		cfg.Redis.Pool.Max = 20
	}
	if cfg.Redis.Pool.Idle == 0 {
		cfg.Redis.Pool.Idle = 10
	}
	if cfg.Redis.Pool.Timeout == 0 {
		cfg.Redis.Pool.Timeout = 300
	}
}

func (cfg *ymlConfig) isCorrect() bool {
	return cfg.checkObj(cfg.Redis) == nil
}
