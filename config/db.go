package config
/**
	数据库配置信息
 */
type dbConfig struct{
	UserName string
	PassWord string
	Host string
	Port int
	DbName string
	DriverName string
}

var (
	//主库
	MasterDbConfig = &dbConfig{
		"root",
		"123456",
		"localhost",
		3306,
		"smy",
		"mysql",
	}
	//从库
	SlaveDbConfig=&dbConfig{
		"root",
		"123456",
		"localhost",
		3306,
		"test",
		"mysql",
	}

)