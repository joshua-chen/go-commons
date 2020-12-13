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
	"github.com/joshua-chen/go-commons/datasource/query"
	"github.com/joshua-chen/go-commons/mvc/context/request"

)

//
// 操作/动作
type Action struct {
	ID          int64  `xorm:"pk autoincr bigint notnull" json:"id"`
	Name        string `xorm:"varchar(50) notnull" json:"name"`
	Type        string `xorm:"varchar(50) notnull" json:"type"` //code,api
	DisplayName string `xorm:"varchar(50) notnull" json:"display_name"`
	CreateAt    int64  `xorm:"bigint notnull" json:"create_at"`
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
func GetActionsByMenuID(menu_id int64, page *request.Pagination) ([]*Action, int64, error) {
	e := datasource.MasterEngine()
	sql := `
SELECT * FROM sys_action
WHERE id in
(
SELECT mid FROM sys_menu_act WHERE menu_id=?
) `

	entities := make([]*Action, 0)

	count, err := query.New(e).PaginationSQL(page, sql, menu_id).FindAndCount(&entities)

	return entities, count, err
}
