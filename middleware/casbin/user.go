/*
 * @Descripttion: 
 * @version: 
 * @Author: joshua
 * @Date: 2020-05-28 15:04:13
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 16:40:21
 */ 
package casbin

import (
	"commons/utils/security/aes"
	"commons/datasource"
	"commons/middleware/models"
	"strconv"
	"time"

	"github.com/kataras/golog"

)

const (
	username = "root"
	password = "123456"
)

// 检查超级用户是否存在
func CheckRootExit() bool {
	e := datasource.MasterEngine()
	// root is existed?
	exit, err := e.Exist(&models.User{Username: username})
	if err != nil {
		golog.Fatalf("@@@ When check Root User is exited? happened error. %s", err.Error())
	}
	if exit {
		golog.Info("@@@ Root User is existed.")

		// 初始化rbac_model
		r := models.User{Username: username}
		if exit, _ := e.Get(&r); exit {
			SetRbacModel(strconv.FormatInt(r.Id, 10))
			CreateSystemRole()
		}
	}
	return exit
}


func CreateRoot() {
	newRoot := models.User{
		Username:   username,
		Password:   aes.AESEncrypt([]byte(password)),
		CreateTime: time.Now(),
	}

	e := datasource.MasterEngine()
	if _, err := e.Insert(&newRoot); err != nil {
		golog.Fatalf("@@@ When create Root User happened error. %s", err.Error())
	}
	rooId := strconv.FormatInt(newRoot.Id, 10)
	SetRbacModel(rooId)

	addAllpolicy(rooId)
}

func addAllpolicy(rooId string) {
	// add policy for root
	//p := casbins.GetEnforcer().AddPolicy(utils.FmtRolePrefix(newRoot.Id), "/*", "ANY", ".*")
	e := GetEnforcer()
	p := e.AddPolicy(rooId, "/*", "ANY", ".*", "", "", "", "", "", "超级用户")
	if !p {
		golog.Fatalf("初始化用户[%s]权限失败.", username)
	}

	//
	for _, v := range Components {
		e.AddGroupingPolicy(rooId, v[0])
	}
}


