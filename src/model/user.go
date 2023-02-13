package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Id            int64  `json:"id,omitempty" gorm:"AUTO_INCREMENT" gorm:"primary_key"`
	Name          string `json:"name,omitempty"`
	Password      string `json:"password,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}
