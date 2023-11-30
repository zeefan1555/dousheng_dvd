package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	VideoId uint   `json:"video_id,omitempty"`
	UserId  uint   `json:"user_id,omitempty"`
	Content string `json:"content,omitempty"`
}
