package model

import "github.com/jinzhu/gorm"

type Favorite struct {
	gorm.Model
	UserId  uint `json:"user_id"`
	VideoId uint `json:"video_id"`
	state   int  `json:"state"` //1: 点赞 0: 未点赞
}
