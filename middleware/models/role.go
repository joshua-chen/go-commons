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
	"commons/datasource"
	"commons/mvc/context/request"

)

type CasbinRule struct {
	Id    int64  `xorm:"pk autoincr INT(10) notnull" json:"id" form:"id"`
	PType string `xorm:"varchar(100) index" json:"p_type"`
	Sub    string `xorm:"varchar(100) index" json:"sub"`
	Obj    string `xorm:"varchar(100) index" json:"obj"`
	Act    string `xorm:"varchar(100) index" json:"act"`
	Ext    string `xorm:"varchar(100) index" json:"ex"`
	Name    string `xorm:"varchar(100) index" json:"name"`
	Des    string `xorm:"varchar(100) index" json:"des"`
	CreateTime   string `xorm:"datetime" json:"create_time"`
}
//{"admin", "/admin*", "GET|POST|DELETE|PUT", ".*", "角色管理"},
func(m *CasbinRule) TableName() string {
	return "sys_casbin_rule"
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
