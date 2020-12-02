/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-28 15:44:16
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 15:59:05
 */
package casbin

import (
	"github.com/kataras/golog"

)

var (
	// 定义系统初始的角色
	Components = [][]string{
		{"admin", "/admin*", "GET|POST|DELETE|PUT", ".*", "角色管理"},
		{"demo", "/demo*", "GET|POST|DELETE|PUT", ".*", "demo角色"},
	}
)

// 创建系统默认角色
func CreateSystemRole() bool {
	e := GetEnforcer()

	for _, v := range Components {
		p := e.GetFilteredPolicy(0, v[0])
		if len(p) == 0 {
			if ok := e.AddPolicy(v); !ok {
				golog.Fatalf("初始化角色[%s]权限失败。%s", v)
			}
		}
	}
	return true
}
//获取所有角色
func GetAllRoles() []string {
	e := GetEnforcer()
	roles := e.GetAllRoles()

	return roles
}

 //获取当前命名策略中显示的角色列表
func GetAllNamedRoles() []string {
	e := GetEnforcer()
	allNamedRoles := e.GetAllNamedRoles("g")

	return allNamedRoles
}


