/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-28 21:42:15
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-29 00:29:40
 */
package config

import (
	_ "sync"

	_ "github.com/joshua-chen/go-commons/exception"
	"github.com/joshua-chen/go-commons/utils"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"

)

func RegisterView(app *iris.Application) {

	for _, viewPath := range AppConfig.ViewPath {
		if viewPath != "" {
			registerView(app, viewPath)
		}
	}
}

func registerView(app *iris.Application, viewPath string) bool {
	path := utils.GetAbsolutePath(viewPath)
	exists, _ := utils.PathExisted(path)
	if !exists {
		golog.Warnf("[registerView]==> %s, not exist! register ignored! ", path)
		return false
	}

	app.RegisterView(iris.HTML(path, ".html"))
	golog.Infof("[registerView]==> %s, ok", path)

	return true
}
