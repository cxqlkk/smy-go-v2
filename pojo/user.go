package pojo

import (
	"time"
)

type User struct {
	Id         string    `form:"id" xorm:"not null pk VARCHAR(32) 'id'"`
	UserName   string    `form:"userName" xorm:"not null VARCHAR(32) 'userName'"`
	PhoneNum   string    `form:"phoneNum" xorm:"not null VARCHAR(11) 'phoneNum'"`
	Password   string    `form:"password" xorm:"not null VARCHAR(11) 'password'"`
	Version    int       `form:"version" xorm:"INT(11)  'version'"`
	CreateTime time.Time `xorm:"DATETIME created 'createTime'"`
	UpdateTime time.Time `xorm:"DATETIME updated 'updateTime'"`
	Status     int       `xorm:"INT(11) 'status'"`
}

func(u *User)TableName()string{
	return "tb_user"
}

