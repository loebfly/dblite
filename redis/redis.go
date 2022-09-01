package redis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"net"
	"time"
)

var ctl = new(control)

type control struct {
	redis *redis.Client
}

func (c *control) connect() error {
	addr := fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)
	redisClient := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     config.Redis.Password,
		DB:           config.Redis.Database,
		PoolSize:     config.Redis.Pool.Max,
		MinIdleConns: config.Redis.Pool.Min,
		IdleTimeout:  time.Duration(config.Redis.Pool.Idle) * time.Minute,
		DialTimeout:  time.Duration(config.Redis.Pool.Timeout) * time.Second,
		Dialer: func() (net.Conn, error) {
			netDialer := &net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 5 * time.Minute,
			}
			return netDialer.Dial("tcp", addr)
		},
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		return err
	}
	c.redis = redisClient
	return nil
}

func (c *control) disconnect() {
	if c.redis != nil {
		_ = c.redis.Close()
		c.redis = nil
	}
}

func (c *control) retry() error {
	if !config.isCorrect() {
		return errors.New("retry connect failed, config is not correct")
	}
	if c.redis == nil {
		err := c.connect()
		if err != nil {
			return errors.New("retry connect failed, " + err.Error())
		}
	}
	_, err := c.redis.Ping().Result()
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

func (c *control) getDB() (*redis.Client, error) {
	if c.redis == nil {
		err := c.connect()
		if err != nil {
			return nil, err
		}
	}
	return c.redis, nil
}
