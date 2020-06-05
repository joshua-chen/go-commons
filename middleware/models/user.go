/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-17 00:11:14
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 16:26:34
 */
package models

import (
	"time"

)

type User struct {
	Id         int64     `json:"id"`
	Username   string    `json:"username"`
	Password   string    `xorm:"notnull" json:"password" form:"password"`
	Token      string    `json:"-"`
	Enabled    int       `xorm:"notnull tinyint(1)" json:"enabled" form:"enabled"`
	Appid      string    `xorm:"notnull" json:"appid" form:"appid"`
	Nickname       string    `xorm:"notnull" json:"nickname" form:"nickname"`
	Phone      string    `xorm:"notnull" json:"phone" form:"phone"`
	Mobile      string    `xorm:"notnull" json:"mobile" form:"mobile"`
	QQ      string    `xorm:"notnull" json:"qq" form:"qq"`
	Email      string    `xorm:"notnull" json:"email" form:"email"`
	Userface   string    `xorm:"notnull" json:"userface" form:"userface"`
	CreateTime time.Time `json:"create_time" form:"create_time"`
	UpdateTime time.Time `json:"update_time" form:"update_time"`
}

func(m *User) TableName() string {
	return "sys_user"
}
