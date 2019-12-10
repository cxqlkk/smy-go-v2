package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/iris-contrib/go.uuid"
	"github.com/kataras/iris/core/errors"
	"smyappTwo/pojo"
	"strings"
)

type ModuleDao interface {
	Create(module *pojo.Module) error
	ListByName(moduleName string) []pojo.ModuleInfo
	GetById(id string) *pojo.Module
	Update(module *pojo.Module) error
}

type moduleDao struct {
	engin *xorm.Engine
}

func NewModuleDao(e *xorm.Engine) ModuleDao {
	return &moduleDao{e}
}

func (m *moduleDao) Create(module *pojo.Module) error {
	if module.Id == "" {
		uid, _ := uuid.NewV4()
		module.Id = strings.ReplaceAll(uid.String(), "-", "")
	}
	module.Status = 1
	n, e := m.engin.InsertOne(module)
	if n == 1 &&e == nil {
		return nil
	}
	return errors.New("插入失败")
}

//用到模块名查寻
func (m *moduleDao) ListByName(moduleName string) []pojo.ModuleInfo {
	list := []pojo.ModuleInfo{}
	e := m.engin.Join("left","tb_fileInfo","tb_fileInfo.id=tb_module.moduleIcon and tb_fileInfo.status=1 ").Where("moduleName like ? and tb_module.status=1" , "%"+moduleName+"%").Find(&list)
	if e != nil {
		return nil
	}
	return list
}

func (m *moduleDao) GetById(id string) *pojo.Module {
	return nil
}

func (m *moduleDao) Update(module *pojo.Module) error {
	return nil
}
