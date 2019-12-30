package cache

import (
	"gameLog/config"
	"github.com/gomodule/redigo/redis"

	"time"
)

var (
	Client *redis.Pool
)

func init() {
	cacheConfig := config.Cfg.Cache
	client := &redis.Pool{
		MaxIdle:     10,
		MaxActive:   100,
		IdleTimeout: 5 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", cacheConfig.Host+":"+cacheConfig.Port,
				redis.DialPassword(cacheConfig.Password),
				redis.DialDatabase(int(cacheConfig.DB)),
				redis.DialConnectTimeout(5*time.Second),
				redis.DialReadTimeout(5*time.Second),
				redis.DialWriteTimeout(5*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
	Client = client
	if err := ping(); err != nil {
		panic(err.Error())
	}
}

func ping() error {
	cacheClient := Client.Get()
	_, err := cacheClient.Do("ping")
	return err
}

func Set(key string, value string, ttl int) error {
	cacheClient := Client.Get()
	defer cacheClient.Close()
	_, err := redis.String(cacheClient.Do("SET", key, value, "EX", ttl))
	return err
}

func Get(key string) (string, error) {
	cacheClient := Client.Get()
	defer cacheClient.Close()
	value, err := redis.String(cacheClient.Do("GET", key))
	return value, err
}

func Remember(key string, ttl int, f func() (string, error)) (string, error) {
	cacheClient := Client.Get()
	defer cacheClient.Close()

	value, err := redis.String(cacheClient.Do("GET", key))
	if err == nil {
		return value, nil
	}
	value, err = f()
	if err != nil {
		return "", err
	}

	_, err = redis.String(cacheClient.Do("SET", key, value, "EX", ttl))

	return value, err
}
