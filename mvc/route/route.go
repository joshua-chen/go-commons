/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-22 15:48:01
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-29 10:52:46
 */
package route

import (
	"strings"

	"github.com/joshua-chen/go-commons/config"
	"github.com/kataras/iris/v12"

)

func PartyFunc(app *iris.Application, path string, fn func(router iris.Party)) {
	api := app.Party(config.AppConfig.APIPrefix.Base)
	api.PartyFunc(path, fn)
}
func PartyCommon(app *iris.Application, fn func(router iris.Party)) {
	path := strings.TrimPrefix(config.AppConfig.APIPrefix.Common, config.AppConfig.APIPrefix.Base)
	PartyFunc(app, path, fn)
}

func PartyWap(app *iris.Application, fn func(router iris.Party)) {
	path := strings.TrimPrefix(config.AppConfig.APIPrefix.Wap, config.AppConfig.APIPrefix.Base)
	PartyFunc(app, path, fn)
}

func PartyWeb(app *iris.Application, fn func(router iris.Party)) {
	path := strings.TrimPrefix(config.AppConfig.APIPrefix.Web, config.AppConfig.APIPrefix.Base)
	PartyFunc(app, path, fn)
}
