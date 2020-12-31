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

	goredis "github.com/go-redis/redis"
	"github.com/joshua-chen/go-commons/config"
	//"github.com/kataras/golog"

)

type Redis struct {
	Client goredis.Client
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
func NewClient() (*goredis.Client)  {
	redisConfig := config.AppConfig.Redis
	if !reflect.DeepEqual(redisConfig, config.Redis{}) && redisConfig.Host != "" {

		url := redisConfig.Host + ":" + redisConfig.Port
		client := goredis.NewClient(&goredis.Options{
			Addr:    url,
			Password: "",
			DB:       0,
		}) 

		return client
	}

	return nil
}
