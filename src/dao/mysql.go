package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMysql() (err error) {
	//格式: 用户名:密码@tcp(IP:端口)/数据库?charset=utf8&parseTime=True&loc=Local
	dsn := fmt.Sprintf("root:zz183173@(192.168.2.205:3306)/dousheng_zeefan?charset=utf8mb4&parseTime=True&loc=Local")

	// 连接数据库
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// 验证数据库连接是否成功，若成功，则无异常
	return DB.DB().Ping()
}

func Close() {
	err := DB.Close()
	if err != nil {
		panic(err)
	}
}
