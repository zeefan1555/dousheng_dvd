package main

import (
	"dousheng_MVC/src/dao"
	"dousheng_MVC/src/model"
	"dousheng_MVC/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	dao.DB.Create(&model.User{})
	dao.DB.AutoMigrate(&model.User{})
	dao.DB.AutoMigrate(&model.Video{})

	r := gin.Default()
	routes.InitRouter(r)

	errrun := r.Run()
	if errrun != nil {
		return

	}

}
