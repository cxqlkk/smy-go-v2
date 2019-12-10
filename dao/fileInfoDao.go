package dao

import (
	"github.com/go-xorm/xorm"
	uuid "github.com/iris-contrib/go.uuid"
	"smyappTwo/pojo"
	"strings"
)

type FileInfoDao interface {
	Create(f *pojo.FileInfo)error

}
type fileInfoDao struct {
	engin *xorm.Engine
}

func NewFileInfoDao(e *xorm.Engine)FileInfoDao{

	return &fileInfoDao{e}
}

func (f *fileInfoDao)Create(fileInfo *pojo.FileInfo)error  {
	fileInfo.Status=1
	if fileInfo.Id==""{
		uid, _ := uuid.NewV4()
		fileInfo.Id = strings.ReplaceAll(uid.String(), "-", "")
	}
	_,e:=f.engin.InsertOne(fileInfo)
	return e


}