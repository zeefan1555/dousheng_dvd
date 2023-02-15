package severice

import (
	"dousheng_MVC/src/dao"
	"dousheng_MVC/src/model"
)

/**
 * @Description 增加user
 * @Param
 * @return
 **/
func CreateUser(user *model.User) (err error) {
	err = dao.DB.Create(&user).Error
	// 响应
	if err != nil {
		return err
	}
	return
}

/**
 * @Description 查询user 通过 名字
 * @Param
 * @return
 **/
func GetUserByName(uname string) (user *model.User, err error) {
	user = new(model.User)
	err = dao.DB.Debug().Where("name = ?", uname).Find(&user).Error
	// 响应
	if err != nil {
		return nil, err
	}
	return user, nil
}

/**
 * @Description 查询user 通过 id
 * @Param
 * @return
 **/
func GetUserById(userId uint) (user model.User, err error) {
	//1.数据模型准备
	user = model.User{}
	//2.在users表中查对应user_id的user
	err = dao.DB.Model(&model.User{}).Where("id = ?", userId).Find(&user).Error
	// 响应
	if err != nil { // 找到
		return user, err
	}
	return user, nil // 未找到

}

/**
 * @Description 查询账号密码
 * @Param
 * @return
 **/
func CheckUser(uname string, pwd string) (user *model.User, err error) {
	user = new(model.User)
	err = dao.DB.Debug().Where("name = ? AND password = ?", uname, pwd).Find(&user).Error
	// 响应
	if err != nil {
		return nil, err
	}
	return user, nil
}
