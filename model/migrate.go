package model

func migrate()  {
	//自动迁移模式
	DB.Set("gorm:table_option","charset=utf8mb4").AutoMigrate(&User{}).AutoMigrate(&Product{}).AutoMigrate(&Admin{}).AutoMigrate(&Order{}).AutoMigrate(&Address{}).
		AutoMigrate(&Cart{})

}
