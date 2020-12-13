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

const PrivilegeTypeCodeMenu   = "menu"
const PrivilegeTypeCodeAction   = "action"  

//
// 权限
type  Privilege struct {
	ID  int64 `xorm:"pk autoincr bigint notnull" json:"id"`
	Name string `xorm:"varchar(50) notnull" json:"name"`
	Type string `xorm:"varchar(50) notnull" json:"type"`//menu，action
	RoleID int64 `xorm:"bigint notnull" json:"role_id"`
	PermID int64 `xorm:"bigint notnull" json:"perm_id"`
}

//
func (m *Privilege) TableName() string {
	return "sys_privilege"
}

//
func CreatePrivilege(p ...*Privilege) (int64, error) {
	e := datasource.MasterEngine()
	return e.Insert(p)
}
func GetMenuPrivilegesByRoleID(role_id int64) ([]*Privilege,  error) {
	e := datasource.MasterEngine() 

	entities := make([]*Privilege, 0)

	err := e.Where("role_id = ?", role_id).And("type=?",PrivilegeTypeCodeMenu).Find(&entities)
	//count, err := query.New(e).PaginationSQL(page, sql, menu_id).FindAndCount(&entities)

	return entities, err
}
//
func GetActionPrivilegesByRoleID(role_id int64) ([]*Privilege,  error) {
	e := datasource.MasterEngine() 

	entities := make([]*Privilege, 0)

	err := e.Where("role_id = ?", role_id).And("type=?",PrivilegeTypeCodeAction).Find(&entities)
	//count, err := query.New(e).PaginationSQL(page, sql, menu_id).FindAndCount(&entities)

	return entities, err
}

 
 
func GetPrivilegesByRoleID(role_id string) ([]*Privilege, error)  {
	e := datasource.MasterEngine()  
	entities := make([]*Privilege, 0)
	err := e.Where("role_id=?",role_id).Find(&entities)

	return entities, err 
}

//
func GetPrivilegesByUID(uid string) ([]*Privilege, error)  {
	e := datasource.MasterEngine()
	sql := `
SELECT * FROM sys_privilege
WHERE role_id in
(
SELECT  role_id FROM sys_user_role ur   WHERE ur.user_id=?
) `

	entities := make([]*Privilege, 0)
	err := e.SQL(sql, uid).Find(&entities)

	return entities, err 
}
 
