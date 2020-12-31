/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-18 16:03:38
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 10:52:19
 */
package redis

import (
	"reflect"
	"sync"

	redigoredis "github.com/garyburd/redigo/redis"
	"github.com/joshua-chen/go-commons/config"
	"github.com/kataras/golog"

)

type Redis struct {
	Conn redigoredis.Conn
}

var (
	instance *Redis
	lock     *sync.Mutex = &sync.Mutex{}
)

func Instance() *Redis {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Redis{}
		}
	}
	return instance
}
func NewConn() (redigoredis.Conn, error) {
	redisConfig := config.AppConfig.Redis
	if !reflect.DeepEqual(redisConfig, config.Redis{}) && redisConfig.Host != "" {
		url := redisConfig.Host + ":" + redisConfig.Port
		c, err := redigoredis.Dial("tcp", url)

		if err != nil {
			golog.Warn("redis connect err", err)
		}

		return c, err
	}

	return nil, nil
}
