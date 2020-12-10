package config

import (
	"strings"

	"github.com/joshua-chen/go-commons/utils"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"

)
func HandleUpload(app *iris.Application) bool{
	//注册静态资源
	uploadPath := AppConfig.UploadPath
	if( strings.Trim(uploadPath,"")== "" ){
		return false
	}


	path := utils.GetAbsolutePath(uploadPath)
	existed, _ := utils.PathExisted(path)
	if !existed {
		golog.Warnf("[HandleUpload]==> %s, not exist! register ignored! ", path)
		return false
	}

	staticDir:= AppConfig.Static.Directory
	if(!strings.HasPrefix(uploadPath,staticDir)){
		app.HandleDir(uploadPath, uploadPath,  router.DirOptions {ShowList: true, Gzip: true, IndexName: "index.html"})
	}
	golog.Infof("[HandleUpload]==> %s, ok", path)
	return true
	//app.HandleDir("/manage/static", staticPath[1])
	//app.HandleDir(staticPath[0]+"/images", staticPath[1]+"/images")
}