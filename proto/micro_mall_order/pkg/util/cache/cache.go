package cache

import (
	"crypto/md5"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

const nameSpace = "track-account_"

type Cache struct {
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Set(redisConn redis.Conn, key, value string, expire int32) error {
	if expire < 0 {
		expire = 0
	}
	_, err := redisConn.Do("SETEX", c.BuildKey(key), expire, value)
	defer redisConn.Close()
	return err
}

func (c *Cache) Get(redisConn redis.Conn, key string) (string, error) {
	reply, err := redisConn.Do("GET", c.BuildKey(key))
	rsp, err := redis.String(reply, err)

	defer redisConn.Close()
	return rsp, err
}

func (c *Cache) HGet(redisConn redis.Conn, key, field string) (string, error) {
	reply, err := redisConn.Do("HGET", c.BuildKey(key), field)
	defer redisConn.Close()
	data, err := redis.String(reply, err)

	return data, nil
}

func (c *Cache) HSet(redisConn redis.Conn, key, field, value string) error {
	_, err := redisConn.Do("HSET", c.BuildKey(key), field, value)
	defer redisConn.Close()
	return err
}

func (c *Cache) HDelete(redisConn redis.Conn, key, field string) error {
	_, err := redisConn.Do("HDEL", c.BuildKey(key), field)
	defer redisConn.Close()

	return err
}

func (c *Cache) Delete(redisConn redis.Conn, key string) error {
	defer redisConn.Close()
	_, err := redisConn.Do("DEL", c.BuildKey(key))
	return err
}

func (c *Cache) Exists(redisConn redis.Conn, key string) (bool, error) {
	defer redisConn.Close()

	r, err := redisConn.Do("EXISTS", c.BuildKey(key))
	return redis.Bool(r, err)
}

func (c *Cache) BuildKey(key string) string {
	if len(key) > 32 {
		m := md5.New()
		m.Write([]byte(key))
		key = fmt.Sprintf("%x", m.Sum(nil))
	}

	return nameSpace + key
}
