package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var DB *gorm.DB

func Database(connn string)  {
	db,err:=gorm.Open("mysql",connn)
	if err != nil {
		panic("数据库连接错误")
	}
	fmt.Println("数据库连接成功")
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(20)//设置连接池
	db.DB().SetMaxOpenConns(100)//设置最大连接数
	DB = db
	migrate()
}
