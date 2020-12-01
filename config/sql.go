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
	_ "fmt"
	_ "sync"

	_ "github.com/CloudyKit/jet"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joshua-chen/go-commons/exception"
	"github.com/joshua-chen/go-commons/utils"
	"github.com/kataras/golog"
	_ "github.com/xormplus/core"
	"github.com/xormplus/xorm"

)

func RegisterSql(engine *xorm.Engine) {
	//注册SqlMap配置，可选功能，如应用中无需使用SqlMap，可无需初始化
	//此处使用xml格式的配置，配置文件根目录为"./sql/oracle"，配置文件后缀为".xml"
	path := utils.GetAbsolutePath(AppConfig.SQLPath) 
	golog.Info("SQLPath==========>" + path)
	err := engine.RegisterSqlMap(xorm.Xml(path, ".xml"))
	if err != nil {
		exception.Instance().Fatal(err)
	}
	//注册动态SQL模板配置，可选功能，如应用中无需使用SqlTemplate，可无需初始化
	//此处注册动态SQL模板配置，使用Pongo2模板引擎，配置文件根目录为"./sql/oracle"，配置文件后缀为".stpl"
	err = engine.RegisterSqlTemplate(xorm.Jet(path, ".jet"))
	if err != nil {
		exception.Instance().Fatal(err)
	}

	//开启SqlMap配置文件和SqlTemplate配置文件更新监控功能，将配置文件更新内容实时更新到内存，如无需要可以不调用该方法
	err = engine.StartFSWatcher()
	if err != nil {
		exception.Instance().Fatal(err)
	}
}
