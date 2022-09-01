package mongo

import (
	"errors"
	"github.com/loebfly/dblite/yml"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var config = new(ymlConfig)

type ymlConfig struct {
	Mongo yml.Mongo `yaml:"mongo"`
}

func (cfg *ymlConfig) Init(ymlPath string) error {
	file, err := ioutil.ReadFile(ymlPath)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(file, cfg); err != nil {
		return err
	}
	if err := cfg.checkObj(cfg.Mongo); err != nil {
		return err
	}
	cfg.fillNull()
	return nil
}

func (cfg *ymlConfig) InitObj(obj yml.Mongo) error {
	if err := cfg.checkObj(obj); err != nil {
		return err
	}
	cfg.Mongo = obj
	cfg.fillNull()
	return nil
}

func (cfg *ymlConfig) checkObj(obj yml.Mongo) error {
	if obj.Url == "" {
		return errors.New("mongo.Url not null")
	}
	if obj.Database == "" {
		return errors.New("mongo.database not null")
	}
	return nil
}

func (cfg *ymlConfig) fillNull() {
	if cfg.Mongo.PoolMax == 0 {
		cfg.Mongo.PoolMax = 20
	}
}

func (cfg *ymlConfig) isCorrect() bool {
	return cfg.checkObj(cfg.Mongo) == nil
}
