package routes

import (
	"dousheng_MVC/src/controller"
	"dousheng_MVC/src/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// -----basic apis---基础接口
	//feed
	apiRouter.GET("/feed/", controller.Feed)

	//user路由组
	apiRouter.GET("/user/", middleware.JwtMiddleware(), controller.UserInfo) //获取用户信息需要验证token
	apiRouter.POST("/user/register/", controller.Register)                   //需要颁发token
	apiRouter.POST("/user/login/", controller.Login)                         //需要颁发token

	////publish路由组
	//apiRouter.POST("/publish/action/", controller.Publish)
	//apiRouter.GET("/publish/list/", controller.PublishList)

	//-----------------

	// extra apis - I
	//apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	//apiRouter.GET("/favorite/list/", controller.FavoriteList)
	//apiRouter.POST("/comment/action/", controller.CommentAction)
	//apiRouter.GET("/comment/list/", controller.CommentList)

}
