package models

 

/** gov doc
http://www.xorm.io/docs/
*/

type (
	// 菜单表
	MenuAction struct {
		MenuID   int64  `xorm:"bigint notnull" json:"menu_id"`
		ActionID string `xorm:"bigint notnull" json:"action_id"`
	}
)

func (m *MenuAction) TableName() string {
	return "sys_menu_act"
}
