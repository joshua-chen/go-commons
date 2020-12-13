/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-27 15:43:36
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-29 00:28:24
 */
package models

import (
	"github.com/joshua-chen/go-commons/datasource"
	"github.com/joshua-chen/go-commons/datasource/time"
	"github.com/joshua-chen/go-commons/mvc/context/request"

)
type Role struct {
	ID       int64  `xorm:"pk autoincr bigint notnull" json:"id" form:"id"`
	Name      string `xorm:"varchar(100) index" json:"name"`
	Description     string `xorm:"varchar(100) index" json:"description"`
	CreateAt time.JsonTime `xorm:"datetime" json:"create_at"  form:"create_at"`
}
type CasbinRule struct {
	Id         int64  `xorm:"pk autoincr bigint notnull" json:"id" form:"id"`
	PType      string `xorm:"varchar(100) index" json:"p_type"`
	Sub        string `xorm:"varchar(100) index" json:"sub"`
	Obj        string `xorm:"varchar(100) index" json:"obj"`
	Act        string `xorm:"varchar(100) index" json:"act"`
	Suf        string `xorm:"varchar(100) index" json:"suf"`
	Name       string `xorm:"varchar(100) index" json:"name"`
	Des        string `xorm:"varchar(100) index" json:"des"`
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
	return e.ID(role.Id).Update(role)
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

 