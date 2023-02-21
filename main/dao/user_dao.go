package dao

import (
	"tiktok/main/common"
	. "tiktok/main/config"
)

func SelectUserByUsername(username string) (user common.User) {
	DB.Where("name = ?", username).First(&user)
	return user
}
func SelectUserByUserId(userId int64) (user common.User) {
	DB.Where("user_id = ?", userId).First(&user)
	return user
}

func InsertUser(user common.User) (int64, error) {
	result := DB.Create(&user)
	if result.Error != nil {
		return -1, result.Error
	}
	return user.UserId, result.Error
}
