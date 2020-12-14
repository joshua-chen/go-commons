/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-27 15:43:36
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-29 00:28:24
 */
package perm

import (
	"github.com/joshua-chen/go-commons/datasource"
	"github.com/joshua-chen/go-commons/datasource/time"
	"github.com/joshua-chen/go-commons/exception"
	"github.com/joshua-chen/go-commons/mvc/context/request"

)

type Role struct {
	ID          int64         `xorm:"pk autoincr bigint notnull" json:"id" form:"id"`
	Name        string        `xorm:"varchar(100) index" json:"name"`
	Description string        `xorm:"varchar(100) index" json:"description"`
	CreateAt    time.JsonTime `xorm:"datetime" json:"create_at"  form:"create_at"`
}
type CasbinRule struct {
	ID       int64  `xorm:"pk autoincr bigint notnull" json:"id" form:"id"`
	PType    string `xorm:"varchar(100) index" json:"p_type"`
	Sub      string `xorm:"varchar(100) index" json:"sub"`
	Obj      string `xorm:"varchar(100) index" json:"obj"`
	Act      string `xorm:"varchar(100) index" json:"act"`
	Suf      string `xorm:"varchar(100) index" json:"suf"`
	Name     string `xorm:"varchar(100) index" json:"name"`
	Des      string `xorm:"varchar(100) index" json:"des"`
	CreateAt string `xorm:"datetime" json:"create_at"  form:"create_at"`
}

//{"admin", "/admin*", "GET|POST|DELETE|PUT", ".*", "角色管理"},
func (m *CasbinRule) TableName() string {
	return "sys_privilege_casbin"
}

func GetPaginationRoles(page *request.Pagination) ([]*CasbinRule, int64, error) {
	e := datasource.MasterEngine()
	roleList := make([]*CasbinRule, 0)

	s := e.Where("p_type=?", "p").Limit(page.Limit, page.Offset)
	if page.SortName != "" {
		switch page.SortOrder {
		case "asc":
			s.Asc(page.SortName)
		case "desc":
			s.Desc(page.SortName)
		}
	}
	count, err := s.FindAndCount(&roleList)

	return roleList, count, err
}

func UpdateRoleById(role *CasbinRule) (int64, error) {
	e := datasource.MasterEngine()
	return e.ID(role.ID).Update(role)
}

func DeleteByRoles(rids []int64) (effect int64, err error) {
	e := datasource.MasterEngine()

	cr := new(CasbinRule)
	for _, v := range rids {
		i, err1 := e.ID(v).Delete(cr)
		effect += i
		err = err1
	}
	return
}

//
func GetAllPrivilegedRoles(privilege_type string) ([]*Role, error) {
	e := datasource.MasterEngine()
	sql := `
SELECT * FROM sys_role
WHERE id in
(
SELECT DISTINCT role_id FROM sys_privilege WHERE type=?
) `

	entities := make([]*Role, 0)
	err := e.SQL(sql, privilege_type).Find(&entities)

	return entities, err
}



// 通过uid获取用户的所有角色
func GetAllRolesByUID(uid int64) ([]*Role, error) {
	e := datasource.MasterEngine()
	sql := `
SELECT * FROM sys_role
WHERE id in
(
SELECT  role_id FROM sys_user_role ur  WHERE ur.user_id=?
) `

	entities := make([]*Role, 0)
	err := e.SQL(sql, uid).Find(&entities)

	return entities, err 
}

//获取所有角色
func GetAllRoles() []*Role {
	e := datasource.MasterEngine()
	entities := make([]*Role, 0)
	err := e.Find(&entities)

	if err != nil {
		exception.Fatal(err)
	}

	return entities
}
