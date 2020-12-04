/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-28 15:43:35
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 16:40:58
 */
package models

import (
	"github.com/joshua-chen/go-commons/datasource"

)

// 角色-菜单关联表
type RoleMenu struct {
	Id  int64 `xorm:"pk autoincr bigint notnull" json:"id"`
	RoleId int64 `xorm:"pk autoincr bigint notnull" json:"role_id"`
	Mid int64 `xorm:"pk autoincr bigint notnull" json:"mid"`
}

func (m *RoleMenu) TableName() string {
	return "sys_role_menu"
}

//
func CreateRoleMenu(roleMenu ...*RoleMenu) (int64, error) {
	e := datasource.MasterEngine()
	return e.Insert(roleMenu)
}
