package main

import (
	"list/dao"
	"list/models"
	"list/routers"

	//_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
)

func main() {

	//创建数据库

	//连接数据库

	dao.Connect()

	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	//启动router
	routers.SetupRouter()
}
