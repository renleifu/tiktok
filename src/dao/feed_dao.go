package dao

import (
	"os"
	"strconv"
	"tiktok/src/common"
	"tiktok/src/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	db  = NewDB()
	dns = config.AppConfig.GetString("datasource.dataSourceName") + "?parseTime=true" //gorm框架 连接mysql
)

type FeedDao struct {
}

func NewDB() *gorm.DB {
	open, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dns,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		println("create gorm filed,err: " + err.Error())
		os.Exit(1)
	}
	return open
}

// SelectVideoByUpdate 根据时间戳获取最多number条视频记录
func (v *FeedDao) SelectVideoByUpdate(unix string, number int) ([]common.Video, error) {
	var video []common.Video
	if len(unix) > 0 {
		// 时间戳转换为时间
		parseInt, err := strconv.ParseInt(unix, 10, 64)
		if err != nil {
			return nil, err
		}
		date := time.UnixMilli(parseInt).Format("2006-01-02 15:04:05")
		db.Preload("User").Limit(number).Where("update_date < ?", date).Order("update_date DESC").Find(&video)
		return video, nil
	}
	db.Preload("TbUser").Limit(number).Order("update_date DESC").Find(&video)
	return video, nil
}

// 根据用户和视频Id查询是否有点赞
func (v *FeedDao) IsFavorite(userID int64, videoID int64) bool {
	var count int64
	db.Model(&common.Favorite{}).Where("`user_id` = ? and `video_id` = ?", userID, videoID).Count(&count)
	return count > 0
}

// 根据用户名id查询是否有关注 true 表示关注
func (v *FeedDao) IsFollwer(originUserID int64, targetUserID int64) bool {
	var size int64
	db.Model(&common.Friendship{}).Where("`follower_id` = ? and `following_id` = ?", originUserID, targetUserID).Count(&size)
	return size > 0
}
