/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-27 14:28:19
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 11:05:19
 */
package middleware

import (
	_ "commons/config"
	"commons/middleware/auth"
	 "commons/middleware/casbin"
	 "commons/middleware/recover"
	"commons/middleware/cors"
	"commons/middleware/jwt"
	_ "commons/utils"
	_ "strings"
	"sync"

	_ "github.com/kataras/golog"
	_ "github.com/kataras/iris"
	_ "github.com/kataras/iris/v12/context"

)

var (
	instance *Middleware
	lock     *sync.Mutex = &sync.Mutex{}
)

type Middleware struct {
	Auth   *auth.Auth
	JWT    *jwt.JWT
	Cors   *cors.Cors
	Casbin *casbin.Casbin
	Recover  *recover.Recover
}

func Instance() *Middleware {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Middleware{
				Auth:   auth.Instance(),
				JWT:    jwt.Instance(),
				Cors:   cors.Instance(),
				Casbin: casbin.Instance(),
				Recover: recover.Instance(),
			}
		}
	}
	return instance
}
