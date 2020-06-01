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
	Name       string    `xorm:"notnull" json:"name" form:"name"`
	Phone      string    `xorm:"notnull" json:"phone" form:"phone"`
	Email      string    `xorm:"notnull" json:"email" form:"email"`
	Userface   string    `xorm:"notnull" json:"userface" form:"userface"`
	CreateTime time.Time `json:"createTime" form:"createTime"`
	UpdateTime time.Time `json:"updateTime" form:"updateTime"`
}

func(m *User) TableName() string {
	return "sys_user"
}
