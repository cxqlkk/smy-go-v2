package pojo

import (
	"time"
)

type FileInfo struct {
	Id         string    `xorm:"not null pk VARCHAR(32) 'id'"`
	StoreName  string    `xorm:"VARCHAR(32) 'storeName'"`
	OriginName string    `xorm:"VARCHAR(32) 'originName'"`
	CreateTime time.Time `xorm:"DATETIME created 'createTime'"`
	UpdateTime time.Time `xorm:"DATETIME updated 'updateTime'"`
	Status     int       `xorm:"not null INT(11) 'status'"`
}


func (f *FileInfo)TableName()string{
	return "tb_fileInfo"
}


