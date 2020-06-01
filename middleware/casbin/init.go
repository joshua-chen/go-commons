/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-28 15:38:33
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 16:00:47
 */
package casbin

import (
	_ "github.com/joshua-chen/go-commons/config"
	"github.com/joshua-chen/go-commons/middleware/models"

	"github.com/kataras/golog"
)

func init() {
	initRootUser()
}

func initRootUser() {
	// root is existed?
	if CheckRootExit() {
		return
	}

	// create root user
	CreateRoot()

	ok := CreateSystemRole()
	if ok {
		addRoleMenu()
	}

}

func addRoleMenu() {
	// 添加role-menu关系
	rMenus := []*models.RoleMenu{
		{Rid: 68, Mid: 2},
		{Rid: 68, Mid: 3},
		{Rid: 68, Mid: 4},
		{Rid: 68, Mid: 5},
	}
	effect, err := models.CreateRoleMenu(rMenus...)
	if err != nil {
		golog.Fatalf("**@@@@@@@@@@@0> %d, %s", effect, err.Error())
	}
	golog.Infof("@@@@@@@@@-> %s, %s", effect, err.Error())
}
