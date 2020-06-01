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
	"go-commons/config"

	"github.com/kataras/iris/v12"
)

func PartyFunc(app *iris.Application, path string, fn func(router iris.Party)) {
	api := app.Party(config.AppConfig.ApiPrefix.Base)
	api.PartyFunc(path, fn)
}
func PartyCommon(app *iris.Application, fn func(router iris.Party)) {
	PartyFunc(app, config.AppConfig.ApiPrefix.Common, fn)
}

func PartyWap(app *iris.Application, fn func(router iris.Party)) {
	//api := app.Party(config.AppConfig.ApiPrefix.Wap)
	PartyFunc(app, config.AppConfig.ApiPrefix.Wap, fn)
}

func PartyWeb(app *iris.Application, fn func(router iris.Party)) {
	PartyFunc(app, config.AppConfig.ApiPrefix.Web, fn)
}
