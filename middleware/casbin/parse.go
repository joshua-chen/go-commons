/*
 * @Descripttion: 
 * @version: 
 * @Author: joshua
 * @Date: 2020-05-27 14:42:52
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 11:04:21
 */ 
package casbin

import (
	"github.com/kataras/golog"

)

/**
用于解析权限
 */

// 通过uid获取用户的所有资源
func GetAllResourcesByUID(uid string) map[string]interface{} {
	allRes := make(map[string]interface{})

	e := GetEnforcer()

	myRes := e.GetPermissionsForUser(uid)
	golog.Infof("GetPermissionsForUser=> %s", myRes)

	// 获取用户的隐形角色
	implicitRoles := e.GetImplicitRolesForUser(uid)
	for _, v := range implicitRoles{
		// 查询用户隐形角色的资源权限
		subRes := e.GetPermissionsForUser(v)
		golog.Infof("-------------------------------------------------")
		golog.Infof("subRes[%s], len(res)=> %d", v, len(subRes))
		golog.Infof("subRes[%s], res=> %s", v, subRes)
		golog.Infof("-------------------------------------------------")
		allRes[v] = subRes
	}

	allRes["myRes"] = myRes
	return allRes
}

// 通过uid获取用户的所有角色
func GetAllRoleByUID(uid string) []string {
	e := GetEnforcer()
	roles := e.GetImplicitRolesForUser(uid)
	
	golog.Infof("roles=> %s", roles)
	return roles
}


func GetPermissionsByUID(uid string) [][]string {
	e := GetEnforcer()
	perms := e.GetPermissionsForUser(uid)
	
	golog.Infof("perms=> %s", perms)
	return perms
}

func AddPermissionForUser(uid string,permission ...string) bool {
	e := GetEnforcer()
	success := e.AddPermissionForUser(uid,permission...)
	
	golog.Infof("AddPermissionForUser.success=> %s", success)
	return success
}

func AddRoleForUser(uid string,role string) bool {
	e := GetEnforcer()
	success := e.AddRoleForUser(uid,role )
	
	golog.Infof("AddRoleForUser.success=> %s", success)
	return success
}