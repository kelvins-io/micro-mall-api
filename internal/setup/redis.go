package setup

import (
	"fmt"
	"gitee.com/cristiane/micro-mall-api/config/setting"
	"github.com/gomodule/redigo/redis"
	"time"
)

func NewRedis(redisSetting *setting.RedisSettingS) (*redis.Pool, error) {
	if redisSetting == nil {
		return nil, fmt.Errorf("RedisSetting is nil")
	}
	if redisSetting.Host == "" {
		return nil, fmt.Errorf("Lack of redisSetting.Host")
	}
	if redisSetting.Password == "" {
		return nil, fmt.Errorf("Lack of redisSetting.Password")
	}
	if redisSetting.PoolNum <= 0 {
		return nil, fmt.Errorf("Wrong redisSetting.PoolNum config")
	}

	return &redis.Pool{
		MaxIdle:     redisSetting.PoolNum,
		MaxActive:   redisSetting.PoolNum,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisSetting.Host)
			if err != nil {
				return nil, err
			}
			if redisSetting.Password != "" {
				if _, err := c.Do("AUTH", redisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if redisSetting.DB != 0 {
				_, err = c.Do("SELECT", redisSetting.DB)
				if err != nil {
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}, nil
}
