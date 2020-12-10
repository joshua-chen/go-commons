package config

import (
	"reflect"
	"strings"

	"github.com/joshua-chen/go-commons/utils"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"

)

func HandleStatic(app *iris.Application) bool {
	//注册静态资源
	static := AppConfig.Static
	if (reflect.DeepEqual(static, Static{}) || strings.Trim(static.Directory, "") == "") {
		return false
	}
	path := utils.GetAbsolutePath(static.Directory)
	existed, _ := utils.PathExisted(path)
	if !existed {
		golog.Warnf("[HandleStatic]==> %s, not exist! register ignored! ", path)
		return false
	}

	//api.HandleDir("/static", "./assets",  DirOptions {ShowList: true, Gzip: true, IndexName: "index.html"})
	app.HandleDir(static.RequestPath, static.Directory, router.DirOptions{ShowList: true, Gzip: true, IndexName: "index.html"})
	golog.Infof("[HandleStatic]==> %s, ok", path)
	return true
	//app.HandleDir("/manage/static", staticPath[1])
	//app.HandleDir(staticPath[0]+"/images", staticPath[1]+"/images")
}
