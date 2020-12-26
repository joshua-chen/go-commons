package perm

import (
	"strconv"

	"github.com/joshua-chen/go-commons/datasource"
	"github.com/joshua-chen/go-commons/datasource/query"
	"github.com/joshua-chen/go-commons/datasource/time"
	"github.com/joshua-chen/go-commons/mvc/context/request"

)

/** gov doc
http://www.xorm.io/docs/
*/

type (
	// 菜单表
	Menu struct {
		ID          int64     `xorm:"pk autoincr bigint notnull" json:"id"`
		URL         string    `xorm:"varchar(255) notnull" json:"url"`
		Redirect     string    `xorm:"varchar(255) notnull" json:"redirect"`
		Component   string    `xorm:"varchar(64) notnull" json:"component"`
		Name        string    `xorm:"varchar(64) notnull" json:"name"`
		Title  string    `xorm:"varchar(64) notnull" json:"title"`
		Icon        string    `xorm:"varchar(64) notnull" json:"icon"`
		KeepAlive   string    `xorm:"varchar(64) notnull" json:"keep_alive"`
		RequireAuth string    `xorm:"varchar(64) notnull" json:"require_auth"`
		ParentID   string    `xorm:"bigint notnull" json:"parent_id"`
		Hidden     bool    `xorm:"tinyint(1) notnull" json:"hidden"`
		Enabled     string    `xorm:"tinyint(1) notnull" json:"enabled"`
		CreateAt  time.JsonTime `json:"create_at"`
		UpdateAt  time.JsonTime `json:"update_at"`

		Children []*Menu `xorm:"-" json:"children"`
		Level        int    `xorm:"-" json:"level"`
		LevelPath        string    `xorm:"-" json:"level_path"`
	}

)

func (m *Menu) TableName() string {
	return "sys_menu"
}



func GetMenusByRoleID(role_id int64, page *request.Pagination) ([]*Action, int64, error) {
	e := datasource.MasterEngine()
	sql := `
SELECT * FROM sys_menu
WHERE id in
(
SELECT obj_id FROM sys_privilege WHERE  role_id=? and type=?
) `

	entities := make([]*Action, 0)

	count, err := query.New(e).PaginationSQL(page, sql, role_id, PrivilegeTypeCodeMenu).FindAndCount(&entities)

	return entities, count, err
}


//
func GetNestedMenus(menus []*Menu) []*Menu {
	out := make([]*Menu, 0)
	if menus == nil || len(menus) == 0 {
		return out
	}

	var parent *Menu

	for _, node := range menus {
		if node.ParentID == "" {
			out = append(out, node)
		} else {
			id,_:= strconv.ParseInt( node.ParentID,10,0)
			parent = GetLocalMenuByID(id, menus)
			if parent.Children != nil {
				//node["parent_index"] = parent.parent_index + "." + (parent.children.length + 1);
				parent.Children = append(parent.Children, node)
			} else {
				//node["parent_index"] = parent.parent_index + "." + 1;
				parent.Children = make([]*Menu, 0)
				parent.Children = append(parent.Children, node)
			}
		}
	}

	return out
}
//
func GetLocalMenuByID(id int64, localMenus []*Menu) *Menu {
	for i := 0; i < len(localMenus); i++ {
		node := localMenus[i]
		if node.ID == id {
			return node
		}
	}
	return nil
}