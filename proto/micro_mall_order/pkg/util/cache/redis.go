package cache

import (
	"crypto/md5"
	"fmt"
	"gitee.com/cristiane/micro-mall-order/vars"
	"github.com/gomodule/redigo/redis"
)

func Set(pool *redis.Pool, key, value string, expire int32) error {
	redisCon := pool.Get()
	_, err := redisCon.Do("SETEX", buildRedisKey(key), expire, value)
	defer redisCon.Close()
	return err
}

func SetNx(pool *redis.Pool, key, value string, expire int32) error {
	if expire < 0 {
		expire = 0
	}
	redisCon := pool.Get()
	_, err := redisCon.Do("SETNX", buildRedisKey(key), expire, value)
	defer redisCon.Close()
	return err
}

func SetEx(pool *redis.Pool, key, value string, expire int32) error {
	if expire < 0 {
		expire = 0
	}
	redisCon := pool.Get()
	_, err := redisCon.Do("SETTEX", buildRedisKey(key), expire, value)
	defer redisCon.Close()
	return err
}

func Get(pool *redis.Pool, key string) (string, error) {
	redisCon := pool.Get()
	reply, err := redisCon.Do("GET", buildRedisKey(key))
	rsp, err := redis.String(reply, err)

	defer redisCon.Close()
	return rsp, err
}

func HGet(pool *redis.Pool, key, field string) (string, error) {
	redisCon := pool.Get()
	reply, err := redisCon.Do("HGET", buildRedisKey(key), field)
	defer redisCon.Close()
	data, err := redis.String(reply, err)

	return data, nil
}

func HSet(pool *redis.Pool, key, field, value string) error {
	redisCon := pool.Get()
	_, err := redisCon.Do("HSET", buildRedisKey(key), field, value)
	defer redisCon.Close()
	return err
}

func HDelete(pool *redis.Pool, key, field string) error {
	redisCon := pool.Get()
	_, err := redisCon.Do("HDEL", buildRedisKey(key), field)
	defer redisCon.Close()

	return err
}

func Delete(pool *redis.Pool, key string) error {
	redisCon := pool.Get()
	defer redisCon.Close()
	_, err := redisCon.Do("DEL", buildRedisKey(key))
	return err
}

func Exists(pool *redis.Pool, key string) (bool, error) {
	redisCon := pool.Get()
	defer redisCon.Close()

	r, err := redisCon.Do("EXISTS", buildRedisKey(key))
	return redis.Bool(r, err)
}

func buildRedisKey(key string) string {
	if len(key) > 32 {
		m := md5.New()
		m.Write([]byte(key))
		key = fmt.Sprintf("%x", m.Sum(nil))
	}

	return vars.AppName + "-" + key
}
