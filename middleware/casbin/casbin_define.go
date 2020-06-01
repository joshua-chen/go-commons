/*
 * @Descripttion: 
 * @version: 
 * @Author: joshua
 * @Date: 2020-05-27 14:46:37
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-27 15:39:44
 */ 
package casbin

// 前端请求的结构体
type (
	RoleDefine struct {
		// 角色的标识等于casbin的sub，但角色需要加role_前缀
		Sub string `json:"sub"`
		// 对应casbin model的定义
		Obj string `json:"obj"`
		Act string `json:"act"`
		Suf string `json:"suf"`
		RoleName string `json:"roleName"`
	}

	// 用户所属角色组
	GroupDefine struct {
		Uid int64    `json:"uid"`
		Sub []string `json:"sub"`
	}
)
