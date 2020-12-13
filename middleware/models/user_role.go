package models

 

/** gov doc
http://www.xorm.io/docs/
*/

type (
	// 用户角色关系表
	UserRole struct {
		RoleID   int64  `xorm:"bigint notnull" json:"role_id"`
		UserID string `xorm:"bigint notnull" json:"user_id"`
	}
)

func (m *UserRole) TableName() string {
	return "sys_user_role"
}
