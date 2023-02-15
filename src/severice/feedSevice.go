package severice

import (
	"dousheng_MVC/src/dao"
	"dousheng_MVC/src/model"
)

func GetFeed() (feedList []model.Video, err error) {
	//获取视频
	err = dao.DB.Debug().Order("created_at DESC").Limit(30).Find(&feedList).Error
	if err != nil {
		return nil, err
	}
	//根据video的id获取作者信息
	for _, feed := range feedList {
		user := model.User{}
		dao.DB.Debug().Model(&user).Find(&user, "id = ?", feed.UserId)
		user.Password = ""
		feed.UserId = user.Id
	}

	return feedList, nil
}
