package dao

import (
	"fmt"
	"github.com/iris-contrib/go.uuid"
	"log"
	"smyappTwo/pojo"
	"testing"
)

func TestCreated(t*testing.T){
	ud:=NewUserDao(MasterEngin())
	fmt.Println(ud)
	u:=&pojo.User{Id:"11111122",UserName:"测试",Status:1}
	e:=ud.Create(u)
	if e!=nil{
		log.Fatal(e)
	}
}


func TestGetById(t*testing.T){
	ud:=NewUserDao(MasterEngin())
	ur:=ud.GetById("123")
	fmt.Println(ur)
}

func TestList(t*testing.T){

	ud:=NewUserDao(MasterEngin())
	ur:=&pojo.User{Status:1}
	usrs:=ud.List(ur,0,10)
	fmt.Println(usrs)
	fmt.Println(ud.Count(ur))
}

func TestUpdate(t*testing.T){
	ud:=NewUserDao(MasterEngin())
	ur:=&pojo.User{Id:"2222",UserName:"updateTest",Version:3}
	e:=ud.Update(ur)
	if e!=nil{
		fmt.Println(e)
	}
}

func TestDelte(t *testing.T){
	ud:=NewUserDao(MasterEngin())
	e:=ud.Delete("11111111")
	if e!=nil{fmt.Println(e)}
}

func TestUUid(t*testing.T){
	u,_:=uuid.NewV4()
	fmt.Println(u.String())
}