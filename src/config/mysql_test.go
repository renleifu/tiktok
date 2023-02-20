package config

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

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

var user User

func Test_init(t *testing.T) {

}

func TestMysqlConnection(t *testing.T) {

	log.Println("连接数据库")
	rows, err := DB.Query("select * from user")
	defer rows.Close()
	if err != nil {
		fmt.Printf("insert data error: #{err}\n")
	}

	for rows.Next() {
		e := rows.Scan(&user.UserId, &user.Name, &user.FollowCount, &user.FollowerCount, &user.IsFollow, &user.Avatar, &user.BackgroundImage, &user.Signature, &user.TotalFavorited, &user.WorkCount, &user.FavoriteCount, &user.Password)
		log.Println("e:", e)
		if e == nil {
			buf, err := json.Marshal(user)
			if err != nil {
				panic(err)
			}
			log.Println(string(buf))
			log.Println(user.UserId, user.Avatar)
		}
	}
	log.Println("连接完成")
}
