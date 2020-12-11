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
	utilspath	"github.com/joshua-chen/go-commons/utils/path"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"

)

//
func RegisterView(app *iris.Application) {

	for _, viewPath := range AppConfig.ViewPath {
		if viewPath != "" {
			registerView(app, viewPath)
		}
	}
}

func registerView(app *iris.Application, viewPath string) bool {
	path := utilspath.GetFullPath(viewPath)
	existed:= utilspath.PathExisted(path)
	if !existed {
		golog.Warnf("[registerView]==> %s, not exist! register ignored! ", path)
		return false
	}

	app.RegisterView(iris.HTML(path, ".html"))
	golog.Infof("[registerView]==> %s, ok", path)

	return true
}
