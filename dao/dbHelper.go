package dao

/**
懒汉模式 实现单例
*/
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	. "smyappTwo/config"
	"sync"
	"xorm.io/core"
)

var (
	masterEngin *xorm.Engine
	oneceMaster sync.Once
	slaveEngin  *xorm.Engine
	oneceSlave  sync.Once
)
//主库
func MasterEngin() *xorm.Engine {
	if masterEngin == nil {
		oneceMaster.Do(func() {
			var err error
			var c = MasterDbConfig
			dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", c.UserName, c.PassWord, c.Host, c.Port, c.DbName)
			masterEngin, err = xorm.NewEngine(c.DriverName, dataSourceName)
			fmt.Println(masterEngin, "masterENgin")
			if err != nil {
				log.Fatal("masterEngin error=", err)
			}
			//	tbMapper:=core.NewPrefixMapper(core.SameMapper{},"tb_")
			//	masterEngin.SetTableMapper(tbMapper)
			masterEngin.ShowSQL(true)
			masterEngin.SetLogLevel(core.LOG_DEBUG)
			masterEngin.SetMaxOpenConns(10)

		})
	}
	return masterEngin
}

//从库
func SlaveEngin() *xorm.Engine {
	if slaveEngin == nil {
		var err error
		var c = SlaveDbConfig
		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", c.UserName, c.PassWord, c.Host, c.Port, c.DbName)
		slaveEngin, err = xorm.NewEngine(c.DriverName, dataSourceName) //此处踩了大坑  slave,err:=  覆盖了全局变量，导致返回为空
		if err != nil {
			log.Fatal("slaveEngin error=", err)
		}
	//	tbMapper := core.NewPrefixMapper(core.SameMapper{}, "tb_")
	//	slaveEngin.SetTableMapper(tbMapper)
		slaveEngin.ShowSQL(true)
		slaveEngin.SetMaxOpenConns(10)
	}

	return slaveEngin
}
