package mongo

import (
	"errors"
	"gopkg.in/mgo.v2"
	"time"
)

var ctl = new(control)

type control struct {
	session *mgo.Session
	mongo   *mgo.Database
}

func (c *control) connect() error {
	session, err := mgo.Dial(config.Mongo.Url)
	if err != nil {
		return err
	}
	session.SetPoolLimit(config.Mongo.PoolMax)
	session.SetMode(mgo.Monotonic, true)
	c.session = session
	return nil
}

func (c *control) disconnect() {
	if c.session != nil {
		c.session.Close()
		c.session = nil
	}
	if c.mongo != nil {
		c.mongo.Session.Close()
		c.mongo = nil
	}
}

func (c *control) retry() error {
	if !config.isCorrect() {
		return errors.New("retry connect failed, config is not correct")
	}
	if c.mongo == nil {
		err := c.connect()
		if err != nil {
			return errors.New("retry connect failed, " + err.Error())
		}
	}
	err := c.mongo.Session.Ping()
	if err != nil {
		return err
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

func (c *control) getDB() (db *mgo.Database, closeDB func(), err error) {
	if c.mongo == nil {
		err = c.connect()
		if err != nil {
			db = nil
			closeDB = nil
			return
		}
		c.mongo = c.session.Copy().DB(config.Mongo.Database)
	}
	db = c.mongo
	err = nil
	closeDB = func() {
		db.Session.Close()
	}
	return
}
