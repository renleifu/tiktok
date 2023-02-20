package service

import (
	. "tiktok/src/cache"
	. "tiktok/src/common"
	. "tiktok/src/config"
	"tiktok/src/utils"
	"time"
)

type UserLoginAndRegisterResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Register(username string, password string) UserLoginAndRegisterResponse {
	var user User
	user.UserId = 0
	user.Name = username
	row := DB.QueryRow("select user_id from user where name = ?", username)
	row.Scan(&user.UserId)
	if user.Exists() {
		println("用户已存在")
		return UserLoginAndRegisterResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		}
	} else {
		user.Password = password
		result, err := DB.Exec(
			"insert into user(name,password) values (?,?)",
			user.Name, user.Password)
		if err != nil {
			panic("新增数据错误")
		}
		user.UserId, err = result.LastInsertId() //新增数据id
		if err != nil {
			panic("新增数据id错误")
		}
		//保存登陆状态
		token := utils.MD5WithSalt(user.Name)
		RCSet(token, user.UserId, 30*time.Minute)
		return UserLoginAndRegisterResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.UserId,
			Token:    token,
		}
	}
}

func Login(username string, password string) UserLoginAndRegisterResponse {
	row := DB.QueryRow(
		"select user_id,name,password from user where name = ?",
		username)
	var user = User{}
	row.Scan(&user.UserId, &user.Name, &user.Password)

	if user.IsCorrect(password) {
		token := utils.MD5WithSalt(username)
		RCSet(token, user.UserId, 30*time.Minute)
		println("登陆成功")
		return UserLoginAndRegisterResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.UserId,
			Token:    token,
		}
	} else {
		return UserLoginAndRegisterResponse{
			Response: Response{StatusCode: 1, StatusMsg: "用户名或密码错误"},
		}
	}
}
