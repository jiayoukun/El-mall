package conf

import (
	"os"
	"test15/cache"
	"test15/model"
)
//var(
//	AppMode string
//	HttpPort string
//	Db string
//	DbHost string
//	DbPort string
//	DbUser string
//	DbPassWord string
//	DbName string
//)

func Init()  {
	//file,err:=ini.Load("./conf/config.ini")
	//if err != nil {
	//	fmt.Println("配置路径读取错误")
	//}
	//LoadService(file)
	//LoadMysql(file)

	os.Setenv("MYSQL_SDN", "root:123456@/El_mall?charset=utf8mb4&parseTime=true&loc=Local")
	os.Setenv("HttpPort", ":3000")
	os.Setenv("REDIS_DB", "1")
	os.Setenv("REDIS_ADDR", "127.0.0.1:6379")
	os.Setenv("REDIS_PW", "123456")
	path:=os.Getenv("MYSQL_SDN")
	//path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	model.Database(path)
	cache.Redis()

}

//
//func LoadService(file *ini.File)  {
//	AppMode=file.Section("service").Key("AppMode").String()
//	HttpPort=file.Section("service").Key("HttpPort").String()
//}
//
//func LoadMysql(file *ini.File)  {
//	Db=file.Section("mysql").Key("Db").String()
//	DbHost=file.Section("mysql").Key("DbHost").String()
//	DbPort=file.Section("mysql").Key("DbPort").String()
//	DbUser=file.Section("mysql").Key("DbUser").String()
//	DbPassWord=file.Section("mysql").Key("DbPassWord").String()
//	DbName=file.Section("mysql").Key("DbName").String()
//}
