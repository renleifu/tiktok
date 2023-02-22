package service

import (
	"tiktok/src/cache"
	"tiktok/src/config"
	"tiktok/src/dao"
)

var (
	videoUrl = config.AppConfig.GetString("video.videoUrl")
	imageUrl = config.AppConfig.GetString("video.imageUrl")
	feedDao  = &dao.FeedDao{}
)

type FeedData struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	NextTime   int64       `json:"next_time"`
	VideoList  []VideoList `json:"video_list"`
}

type Author struct {
	UserID        int64  `json:"user_id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type VideoList struct {
	ID            int64  `json:"id"`
	UserId        int64  `json:"user_id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

func GetFeed(latestTime string, token string) *FeedData {
	data := &FeedData{}
	var userId int64
	userId = -1

	videos, err := feedDao.SelectVideoByUpdate(latestTime, 30)
	if err != nil {
		return &FeedData{
			StatusCode: 400,
			StatusMsg:  "查找视频失败,err： " + err.Error(),
		}
	}

	if len(token) > 0 {
		get := cache.RCGet(token)
		userId, err = get.Int64()
	}

	for _, v := range videos {
		data.VideoList = append(data.VideoList, VideoList{
			ID:            v.Id,
			UserId:        v.UserId,
			PlayUrl:       getVideoUrl(v.PlayUrl),
			CoverUrl:      getImageUrl(v.CoverUrl),
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			Title:         v.Title,
			IsFavorite:    isFavorite(userId, v.Id),
		})
	}
	data.StatusCode = 200
	return data
}

func getVideoUrl(videoName string) string {
	return videoUrl + videoName
}

func getImageUrl(imageName string) string {
	return imageUrl + imageName
}

func isFollow(uId int64, vId int64) bool {
	if uId == -1 {
		return false
	}
	// 查看该uId用户是否关注了该视频作者
	return feedDao.IsFollwer(uId, vId)
}

func isFavorite(uId int64, vId int64) bool {
	if uId == -1 {
		return false
	}
	// TODO 查看该uId用户是否点赞了vId该视频
	return feedDao.IsFavorite(uId, vId)
}
