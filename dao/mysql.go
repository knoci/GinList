package dao

import (
	"fmt"
	"gopkg.in/ini.v1"
	//"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type MysqlConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	DBname   string `ini:"db"`
}

func LoadConfig() *MysqlConfig {

	//development_通过结构体映射参数
	c := new(MysqlConfig)
	ini.MapTo(c, "config/config.ini")
	fmt.Println(c)

	return c
}

func Connect() {

	c := LoadConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", c.Host, c.User, c.Password, c.DBname, c.Port)
	/* mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",c.User, c.Password, c.Host, c.Port, c.DBname)
	dsn := "root:root1234@tcp(127.0.0.1:13306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	*/
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // connect
	// mysql DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Print("==========连接数据库成功==========\n")
}
