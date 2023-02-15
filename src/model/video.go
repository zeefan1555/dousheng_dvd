package model

import "github.com/jinzhu/gorm"

type Video struct {
	gorm.Model
	Id            int64  `json:"id,omitempty" gorm:"AUTO_INCREMENT;primary_key"`
	UserId        int64  `json:"author,omitempty"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount uint   `json:"favorite_count,omitempty"`
	CommentCount  uint   `json:"comment_count,omitempty"` //int64与*int64的区别是:int64的默认值是0,*int64的默认值是nil
	IsFavorite    uint   `json:"is_favorite,omitempty"`   //*bool与bool的区别是:bool的默认值是false,*bool的默认值是nil
	Title         string `json:"title,omitempty"`
}
