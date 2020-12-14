/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-28 15:43:35
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 16:40:58
 */
package perm

import (
	"github.com/joshua-chen/go-commons/datasource"
	"github.com/joshua-chen/go-commons/datasource/query"
	"github.com/joshua-chen/go-commons/mvc/context/request"

)

const ActionTypeRequestMethod = "request_method"
const ActionTypeActionCode = "action_code"
const ActionTypeAPI = "api"

//
// 操作/动作
type Action struct {
	ID          int64  `xorm:"pk autoincr bigint notnull" json:"id"`
	Name        string `xorm:"varchar(50) notnull" json:"name"`
	Type        string `xorm:"varchar(50) notnull" json:"type"` //action_code,api
	Method      string `xorm:"varchar(100) null" json:"method"`
	DisplayName string `xorm:"varchar(50) null" json:"display_name"`
	CreateAt    int64  `xorm:"datetime notnull" json:"create_at"`
}

//
func (m *Action) TableName() string {
	return "sys_action"
}

//
func CreateAction(act ...*Action) (int64, error) {
	e := datasource.MasterEngine()
	return e.Insert(act)
}

//
func GetPaginationActionsByMenuID(page *request.Pagination, menu_id int64) ([]*Action, int64, error) {
	e := datasource.MasterEngine()
	sql := `
SELECT * FROM sys_action
WHERE id in
(
SELECT perm_id FROM sys_menu_act WHERE menu_id=?
) `

	entities := make([]*Action, 0)

	count, err := query.New(e).PaginationSQL(page, sql, menu_id).FindAndCount(&entities)

	return entities, count, err
}

//
func GetActionsByRoleID(role_id int64) ([]*Action, error) {
	e := datasource.MasterEngine()
	sql := `
	SELECT * FROM sys_action
	WHERE id in
	(
	SELECT perm_id FROM sys_privilege WHERE role_id=? and type=?
	) `
	entities := make([]*Action, 0)

	err := e.SQL(sql, role_id, PrivilegeTypeCodeAction).Find(&entities)

	return entities, err
}

func GetActionsByUID(uid int64) ([]*Action, error) {
	e := datasource.MasterEngine()
	sql := `
	SELECT * FROM sys_action
	WHERE id in
	(
	SELECT perm_id FROM sys_privilege p join sys_user_role ur WHERE p.role_id=ur.role_id and ur.user_id=?
	) `
	entities := make([]*Action, 0)

	err := e.SQL(sql, uid, PrivilegeTypeCodeAction).Find(&entities)

	return entities, err
}

func HasPrivilege(uid int64, path string, method string) (bool, error) {
	e := datasource.MasterEngine()
	sql := `
	SELECT count(1) FROM sys_action 
	WHERE type=? and LOCATE(REPLACE(name, '*', ''),?)>0 and LOCATE(?,method)>0 and id in
	(
	SELECT perm_id FROM sys_privilege p join sys_user_role ur WHERE p.role_id=ur.role_id and ur.user_id=? and p.type=?
	) `

	var count int
	_, err := e.SQL(sql, ActionTypeAPI, path, method, uid, PrivilegeTypeCodeAction).Get(&count)

	return count > 0, err
}
