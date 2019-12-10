package demos

import (
	"fmt"
	"testing"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func TestFrist(t*testing.T){

	master,e:=xorm.NewEngine("mysql","root:123456@/smy?charset=utf8&parseTime=true&loc=Local")
	slave,e:=xorm.NewEngine("mysql","root:123456@/test?charset=utf8&parseTime=true&loc=Local")
fmt.Println(master)
	if e!=nil{
		fmt.Println(e)
	}
	slaves:=[]*xorm.Engine{slave}
	g,e:=xorm.NewEngineGroup(master,slaves,xorm.RandomPolicy())
	//xorm.NewEngine("mysql","root:123456@/test?charset=utf8")
	if e!=nil{
		fmt.Println(e)
	}
	g.ShowSQL(true)
	g.Logger().SetLevel(core.LOG_DEBUG)
	prefixMapper:=core.NewPrefixMapper(core.SameMapper{},"tb_")
	g.SetMapper(prefixMapper)
	g.SetColumnMapper(core.SameMapper{})

}
