package services

import (
	"smyappTwo/dao"
	"smyappTwo/pojo"
)

type ModuleService interface {
	Create(module *pojo.Module) error
	ListByName(moduleName string) []pojo.ModuleInfo
	GetById(id string) *pojo.Module
	Update(module *pojo.Module) error
}

type moduleService struct {
	dao dao.ModuleDao
}

func NewModuleService(d dao.ModuleDao) ModuleService {
	return &moduleService{d}
}

func (s *moduleService) Create(module *pojo.Module) error {

	return s.dao.Create(module)
}
func (s *moduleService) ListByName(moduleName string) []pojo.ModuleInfo {

	return s.dao.ListByName(moduleName)
}
func (s *moduleService) GetById(id string) *pojo.Module {
	return s.dao.GetById(id)
}
func (s *moduleService) Update(module *pojo.Module) error {
	return s.dao.Update(module)
}
