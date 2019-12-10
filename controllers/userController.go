package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"log"
	"smyappTwo/pojo"
	"smyappTwo/services"
	"smyappTwo/vo"
)

type UserController struct {
	Service services.UserService
	Ctx     iris.Context //大写
	Session *sessions.Session
}

func (u *UserController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/getById/{id}", "GetById")
	b.Handle("GET","/list/{userName}","List")
	b.Handle("POST","/add","Add")
	b.Handle("GET","/login","Login")
}

//	GET:getById/id
//	查询单条记录
func (u *UserController) GetById() *pojo.User {
	id := u.Ctx.Params().Get("id")
	return u.Service.GetById(id)
}

//GET:list/xm
func (u *UserController) List() interface{} {

	userName := u.Ctx.Params().Get("userName")

	p:=&vo.Page{}
	err:=u.Ctx.ReadForm(p) //获得问号后的参数
	if err!=nil{
		log.Println(err)
	}

	return u.Service.List(&pojo.User{UserName:userName}, p.PageNo, p.PageSize)
}

func (u*UserController)Add()error{
	usr:=&pojo.User{}
	u.Ctx.ReadForm(usr)
	return u.Service.Create(usr)
}
//login

func (u *UserController)Login()interface{}{//暂时返回interface

//Ctx.Params().Get("phoneNum") 无法获取 问号后的参数,获得路径值
//formValues ,readForm 获取的是 问号后的值，
	usr:=&pojo.User{}
	u.Ctx.ReadForm(usr)
	fmt.Println(usr)
	user:= u.Service.GetByPhoneAndPassword(usr.PhoneNum,usr.Password)
	if user!=nil{
		u.Session.Set("uid",user.Id)
	}
	return user
}

func (c *UserController) logout() {
	c.Session.Destroy()
}

