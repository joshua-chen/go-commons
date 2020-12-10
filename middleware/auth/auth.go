/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-27 14:28:19
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-29 10:43:05
 */
package auth

import (
	_ "strings"
	"sync"

	"github.com/joshua-chen/go-commons/config"
	"github.com/joshua-chen/go-commons/middleware/casbin"
	"github.com/joshua-chen/go-commons/middleware/jwt"
	_ "github.com/joshua-chen/go-commons/mvc/context"
	"github.com/joshua-chen/go-commons/utils"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"

)

type Auth struct {
}

var (
	instance *Auth
	lock     *sync.Mutex = &sync.Mutex{}
)

func Instance() *Auth {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Auth{}
		}
	}
	return instance
}
func (a *Auth) New() context.Handler {

	return New()
}

func New() context.Handler {
	handler := func(ctx context.Context) {
		path := ctx.Path()		
		golog.Debug("request path===> ",path)

		// 过滤静态资源、login接口、首页等...不需要验证
		if checkURL(path) {
			ctx.Next()
			return
		}

		//AJAX进行跨域请求时的预检请求 ，跳过检查
		if(ctx.Method() == iris.MethodOptions){
			ctx.Next()
			return
		}

		// jwt token拦截
		if !jwt.Filter(ctx) {
			return
		}

		// casbin权限拦截
		ok := casbin.Filter(ctx)
		if !ok {
			return
		}

		// Pass to real API
		ctx.Next()
	}

	return handler
}

/**
return
	true:则跳过不需验证，如登录接口等...
	false:需要进一步验证
*/
func checkURL(requestPath string) bool {
	
	requestStaticPath := config.AppConfig.Static.RequestPath
	if utils.HasPrefix(requestPath, requestStaticPath) {
		return true
	} 

	anonymousUrls := config.AppConfig.AnonymousRequest.Urls
	for _, v := range anonymousUrls {
		if requestPath == v {
			return true
		}
	}

	anonymousPrefixes := config.AppConfig.AnonymousRequest.Prefixes
	for _, v := range anonymousPrefixes {
		if utils.HasPrefix(requestPath, v) {
			return true
		}
	}

	anonymousSuffixes := config.AppConfig.AnonymousRequest.Suffixes
	for _, v := range anonymousSuffixes {
		if utils.HasSuffix(requestPath, v) {
			return true
		}
	}

	// strings.Index(requestPath,v)
	return false
}
