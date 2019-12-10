package pojo

import (
	"time"
)

type Module struct {
	Id         string    `json:"id" xorm:"not null pk VARCHAR(32) 'id'"`
	ModuleName string    `json:"moduleName"xorm:"VARCHAR(5) 'moduleName'"`
	ModuleIcon string    `json:"moduleIcon "xorm:"VARCHAR(200) 'moduleIcon'"`
	CreateTime time.Time `json:"createTime" xorm:"DATETIME created 'createTime'"`
	UpdateTime time.Time `json:"updateTime" xorm:"DATETIME updated 'updateTime'"`
	Status     int       `json:"status" xorm:"not null INT(11) 'status'"`
}

func (m *Module)TableName()string{
	return "tb_module"
}
// 和
type ModuleInfo struct {
	Module `xorm:"extends""`
	StoreName string `json:"storeName" xorm:"'storeName'"`  //icon 存储的名字
}