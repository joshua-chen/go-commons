package models

import (
	"fmt"
	"commons/datasource"
	"commons/mvc/context/request"
	"strconv"
	"time"

)

/** gov doc
http://www.xorm.io/docs/
 */

type (
	// 菜单表
	Menu struct {
		Id          int64     `xorm:"pk autoincr INT(10) notnull" json:"id"`
		Path        string    `xorm:"varchar(64) notnull" json:"path"`
		Url         string    `xorm:"varchar(64) notnull" json:"url"`
		Modular     string    `xorm:"varchar(64) notnull" json:"modular"`
		Component   string    `xorm:"varchar(64) notnull" json:"component"`
		Name        string    `xorm:"varchar(64) notnull" json:"name"`
		Icon        string    `xorm:"varchar(64) notnull" json:"icon"`
		KeepAlive   string    `xorm:"varchar(64) notnull" json:"keep_alive"`
		RequireAuth string    `xorm:"varchar(64) notnull" json:"require_auth"`
		ParentId    string    `xorm:"INT(10) notnull" json:"parent_id"`
		Enabled     string    `xorm:"tinyint(1) notnull" json:"enabled"`
		CreateTime  time.Time `json:"createTime"`
		UpdateTime  time.Time `json:"updateTime"`

		Children []Children `xorm:"-" json:"children"`
	}

	// 子菜单
	Children struct {
		Id2          int64  `xorm:"pk autoincr INT(10) notnull" json:"id"`
		Path2        string `xorm:"varchar(64) notnull" json:"path"`
		Modular2     string `xorm:"varchar(64) notnull" json:"modular"`
		Component2   string `xorm:"varchar(64) notnull" json:"component"`
		Name2        string `xorm:"varchar(64) notnull" json:"name"`
		Icon2        string `xorm:"varchar(64) notnull" json:"icon"`
		KeepAlive2   string `xorm:"varchar(64) notnull" json:"keep_alive"`
		RequireAuth2 string `xorm:"varchar(64) notnull" json:"require_auth"`
		ParentId2    string `xorm:"INT(10) notnull" json:"parent_id"`
		Enabled2     string `xorm:"tinyint(1) notnull" json:"enabled"`
	}

	// 菜单树
	MenuTreeGroup struct {
		Menu     `xorm:"extends"`
		Children `xorm:"extends"`
	}
)

func(m *Menu) TableName() string {
	return "sys_menu"
}

// 获取用户的菜单列表
func DynamicMenuTree(uid int64) []Menu {
	e := datasource.MasterEngine()
	sql := fmt.Sprintf(`
SELECT
	m1.id, m1.path, m1.modular, m1.component, m1.icon, m1.name, m1.require_auth,
	m2.id AS id2,
	m2.modular AS modular2,
	m2.component AS component2,
	m2.icon AS icon2,
	m2.keep_alive AS keep_alive2,
	m2.name AS name2,
	m2.path AS path2,
	m2.require_auth AS require_auth2
FROM menu m1, menu m2
WHERE m1.id = m2.parent_id
	AND m1.id != 1
	AND m2.id IN 
(
		SELECT rm.mid
		FROM role_menu rm WHERE rm.rid in
		(
			SELECT id FROM casbin_rule 
			WHERE 
			v2 <> 'ANY' AND 
			v0 in 
			(
				SELECT v1 FROM casbin_rule WHERE v0=%d
			)
		)
)
AND m2.enabled=true order by m1.id, m2.id
`, uid)

	menuTree := make([]MenuTreeGroup, 0)
	e.SQL(sql).Find(&menuTree)

	menus := make([]Menu, 0) // 菜单树

	if len(menuTree) > 0 {
		parentMenu := menuTree[0].Menu    // 父级菜单
		childMenus := make([]Children, 0) // 所有的子菜单
		for _, v := range menuTree {
			childMenus = append(childMenus, v.Children)
		}
		parentMenu.Children = childMenus

		menus = append(menus, parentMenu)
	}
	return menus
}

func GetPaginationMenus(page *request.Pagination) ([]*Menu, int64, error) {
	e := datasource.MasterEngine()
	menuList := make([]*Menu, 0)

	s := e.Limit(page.Limit, page.Offset)
	if page.SortName != "" {
		switch page.SortOrder {
		case "asc":
			s.Asc(page.SortName)
		case "desc":
			s.Desc(page.SortName)
		}
	}
	count, err := s.FindAndCount(&menuList)

	return menuList, count, err
}

func GetMenusByRoleid(rid int64, page *request.Pagination) ([]*Menu, int64, error) {
	e := datasource.MasterEngine()
	sql := fmt.Sprintf(`
SELECT * FROM menu
WHERE id in
(
SELECT mid FROM role_menu WHERE rid=%d
)
`, rid)

	if page.SortName != "" {
		sql += " ORDER BY "
		switch page.SortOrder {
		case "asc":
			sql += page.SortName + " ASC"
		case "desc":
			sql += page.SortName + " DESC"
		}
	}
	sql += " LIMIT " + strconv.Itoa(page.Offset) + ", " + strconv.Itoa(page.Limit)

	menus := make([]*Menu, 0)
	err := e.SQL(sql).Find(&menus)

	return menus, 10, err
}
