package controller

import (
	"dousheng_MVC/src/common"
	"dousheng_MVC/src/middleware"
	"dousheng_MVC/src/model"
	"dousheng_MVC/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]model.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

// 作用: 生成一个唯一的id
var userIdSequence = int64(1)

type UserLoginResponse struct {
	common.Response
	UserId uint   `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	common.Response
	User model.User `json:"user"`
}

func Register(c *gin.Context) {
	// 1. 从请求参数中提取用户名和密码
	username := c.Query("username")
	password := c.Query("password")

	// 2. 校验参数
	if len(username) > 32 {
		c.JSON(http.StatusOK, gin.H{"status_code": 1, "status_msg": "用户名太长"})
		return
	}

	if len(password) > 32 {
		c.JSON(http.StatusOK, gin.H{"status_code": 1, "status_msg": "密码太长"})
		return
	}

	//// 3. 查询用户是否已存在
	//user, err := severice.GetUserByName(username)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"status_code": 1, "status_msg": "查询用户失败"})
	//	return
	//}
	//
	//if user != nil {
	//	c.JSON(http.StatusOK, gin.H{"status_code": 1, "status_msg": "用户已存在"})
	//	return
	//}

	// 4. 创建用户
	newUser := model.User{
		Name:     username,
		Password: password,
	}

	if err := severice.CreateUser(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status_code": 1, "status_msg": "创建用户失败"})
		return
	}

	// 5. 生成访问令牌
	token, err := middleware.CreateToken(newUser.ID, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status_code": 1, "status_msg": "生成访问令牌失败"})
		return
	}

	// 6. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "注册成功",
		"user_id":     newUser.ID,
		"token":       token,
	})
}

//func Register(c *gin.Context) {
//	//1.参数提取
//	username := c.Query("username")
//	password := c.Query("password")
//
//	////3.查询用户是否在数据库
//	//checkuser, err := severice.GetUserByName(username)
//	//if checkuser != nil {
//	//	c.JSON(http.StatusOK, UserLoginResponse{
//	//		Response: common.Response{StatusCode: 1, StatusMsg: "User already exist"},
//	//	})
//	//	return
//	//}
//	//MD5加密
//	//password = common.Md5Encrypt(password)
//	//4.创建用户
//	user := model.User{
//		Name:     username,
//		Password: password,
//	}
//
//	err := severice.CreateUser(&user)
//	if err != nil {
//		c.JSON(http.StatusOK, UserLoginResponse{
//			Response: common.Response{StatusCode: 1, StatusMsg: "Create user fail"},
//		})
//		return
//	}
//	//生成 token
//	token, _ := middleware.CreateToken(user.ID, user.Name)
//
//	//5.返回对象
//	c.JSON(http.StatusOK, UserLoginResponse{
//		Response: common.Response{StatusCode: 0},
//		UserId:   user.ID,
//		Token:    token,
//	})
//}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//参数校验
	if len(password) <= 5 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "Password must be greater than 5"},
		})
		return
	}
	//MD5加密
	password = common.Md5Encrypt(password)
	//查询用户是否在数据库
	user, err := severice.CheckUser(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: common.Response{StatusCode: 0},
		UserId:   user.ID,
	})
}

func UserInfo(c *gin.Context) {
	userid := c.MustGet("userid").(int64)
	user, err := severice.GetUserById(userid)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}
	//清空密码
	user.Password = ""
	c.JSON(http.StatusOK, UserResponse{
		Response: common.Response{StatusCode: 0},
		User:     user,
	})
}
