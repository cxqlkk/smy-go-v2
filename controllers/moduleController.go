package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"smyappTwo/pojo"
	"smyappTwo/services"
)

type ModuleController struct {
	Service services.ModuleService
	Ctx iris.Context
}

func (m *ModuleController)BeforeActivation(b mvc.BeforeActivation)  {
	b.Handle("GET","/list","List")

}

func (m *ModuleController)List()[]pojo.ModuleInfo{
	moduleName:=m.Ctx.URLParam("moduleName") // 可以获得 ？ 后的参数， 和readForm
	fmt.Println(moduleName)
	return m.Service.ListByName(moduleName)

}