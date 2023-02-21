package common

import (
	"time"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type User struct {
	UserId          int64  `gorm:"column:user_id;primaryKey;autoIncrement:true"`
	Name            string `gorm:"column:name"`
	FollowCount     int64  `gorm:"column:follow_count"`
	FollowerCount   int64  `gorm:"column:follower_count"`
	IsFollow        bool   `gorm:"column:is_follow"`
	Avatar          string `gorm:"column:avatar"`
	BackgroundImage string `gorm:"column:background_image"`
	Signature       string `gorm:"column:signature"`
	TotalFavorited  int64  `gorm:"column:total_favorited"`
	WorkCount       int64  `gorm:"column:work_count"`
	FavoriteCount   int64  `gorm:"column:favorite_count"`
	Password        string `gorm:"column:password"`
}

func (*User) TableName() string {
	return "user"
}

type Comment struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement:true"`
	UserId     int64     `gorm:"column: user_id"`
	VideoId    int64     `gorm:"column: video_id"`
	Content    string    `gorm:"column: content"`
	CreateDate time.Time `gorm:"column: create_date"`
}

func (*Comment) TableName() string {
	return "comment"
}

type Friendship struct {
	Id        int64 `gorm:"column:id;primaryKey;autoIncrement:true"`
	UserAId   int64 `gorm:"column:userA_id"`
	UserBId   int64 `gorm:"column:userB_id"`
	MessageId int64 `gorm:"column:message_id"`
}

func (*Friendship) TableName() string {
	return "friendship"
}

type Message struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement:true"`
	ToUserId   int64     `gorm:"column:to_user_id"`
	FromUserId int64     `gorm:"column:from_user_id"`
	Content    int64     `gorm:"column:content"`
	CreateTime time.Time `gorm:"column: create_time"`
}

func (*Message) TableName() string {
	return "message"
}

type RelationshipFollow struct {
	Id         int64 `gorm:"column:id;primaryKey;autoIncrement:true"`
	FollowId   int64 `gorm:"column:follow_id"`
	FollowerId int64 `gorm:"column:follower_id"`
}

func (*RelationshipFollow) TableName() string {
	return "relationship_follow"
}

type Video struct {
	Id            int64  `gorm:"column:id;primaryKey;autoIncrement:true"`
	UserId        int64  `gorm:"column: user_id"`
	PlayUrl       string `gorm:"column: play_url"`
	CoverUrl      string `gorm:"column:cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	IsFavorite    bool   `gorm:"column:is_favorite"`
	Title         string `gorm:"column:title"`
}

type Favorite struct {
	Id     int64 `gorm:"column:id;primaryKey;autoIncrement:true"`
	UserId int64 `gorm:"column: user_id"`
	Video  int64 `gorm:"column: video_id"`
}

func (*Video) TableName() string {
	return "video"
}
func (user *User) IsCorrect(password string) bool {
	return user.Password == password
}
func (user *User) Exists() bool {
	return user.UserId > 0
}
