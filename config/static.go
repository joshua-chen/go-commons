package config

import (
	"strings"

	"github.com/joshua-chen/go-commons/utils"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"

)
func HandleStatic(app *iris.Application) bool{
	//注册静态资源
	staticPath := AppConfig.StaticPath
	if( strings.Trim(staticPath.Directory,"")== "" ){
		return false
	}
	path := utils.GetAbsolutePath(staticPath.Directory)
	exists, _ := utils.PathExisted(path)
	if !exists {
		golog.Warnf("[HandleStatic]==> %s, not exist! register ignored! ", path)
		return false
	}

	//api.HandleDir("/static", "./assets",  DirOptions {ShowList: true, Gzip: true, IndexName: "index.html"})
	app.HandleDir(staticPath.RequestPath, staticPath.Directory,  router.DirOptions {ShowList: true, Gzip: true, IndexName: "index.html"})

	return true
	//app.HandleDir("/manage/static", staticPath[1])
	//app.HandleDir(staticPath[0]+"/images", staticPath[1]+"/images")
}