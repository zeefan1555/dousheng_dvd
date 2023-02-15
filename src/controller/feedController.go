package controller

import (
	"dousheng_MVC/src/common"
	"dousheng_MVC/src/model"
	"dousheng_MVC/src/severice"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	common.Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	//根据投稿时间倒序的视频列表，获取三十个
	feed, err := severice.GetFeed()
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  common.Response{StatusCode: 1, StatusMsg: "query sql error"},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
		fmt.Println("Feed error: ", err)
		return
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: feed,
		NextTime:  time.Now().Unix(),
	})
}
