/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-18 09:21:47
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-29 00:14:14
 */
package sys

import (
	"commons/datasource"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"

)

// 主库，单例
func MasterEngine() *xorm.Engine {

	engine := datasource.MasterEngine()
	tbMapper := core.NewPrefixMapper(core.GonicMapper{}, "sys_")
	engine.SetTableMapper(tbMapper)
	return engine
}

// 从库，单例
func SlaveEngine() *xorm.Engine {
	engine := datasource.SlaveEngine()
	tbMapper := core.NewPrefixMapper(core.GonicMapper{}, "sys_")
	engine.SetTableMapper(tbMapper)
	return engine
}
 
