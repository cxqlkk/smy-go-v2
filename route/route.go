package route

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"smyappTwo/controllers"
	"smyappTwo/dao"
	"smyappTwo/services"
	"time"
)

//var validate  *validator.Validate

func Route(app *iris.Application){
	//validate=validator.New()
	mvc.Configure(app.Party("/user"),UserRoute)
	mvc.Configure(app.Party("/module"),ModuleRoute)


}

func UserRoute( app *mvc.Application){
	engin:=dao.MasterEngin()
	dao:=dao.NewUserDao(engin)
	service:=services.NewUserService(dao)
	sessManger:=sessions.New(sessions.Config{Cookie:"smy",Expires:24*time.Hour})
	app.Register(service,sessManger.Start)
	app.Handle(new(controllers.UserController))


}
//模块管理
func ModuleRoute( app *mvc.Application)  {
	dao:=dao.NewModuleDao(dao.MasterEngin())
	moduleService:=services.NewModuleService(dao)
	app.Register(moduleService)
	app.Handle(new(controllers.ModuleController))
}