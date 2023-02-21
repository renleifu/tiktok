package service

import (
	"golang.org/x/crypto/bcrypt"
	. "tiktok/main/common"
	. "tiktok/main/dao"
	. "tiktok/main/utils"
)

type UserLoginAndRegisterResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}
type UserInfoResponse struct {
	Response
	User User `json:"user"`
}

func Register(username string, password string) UserLoginAndRegisterResponse {
	user := SelectUserByUsername(username)
	if user.Exists() {
		return UserLoginAndRegisterResponse{
			Response: Response{StatusCode: 422, StatusMsg: "User already exist"},
		}
	} else {
		user.Name = username
		user.Password = password
		user.UserId, _ = InsertUser(user)
		if user.UserId == -1 {
			panic("新增数据错误")
		}
		//保存登陆状态
		token, err := CreateToken(user.UserId, user.Name)
		if err != nil {
			return UserLoginAndRegisterResponse{
				Response: Response{StatusCode: 500, StatusMsg: "服务器错误"},
			}
		}
		return UserLoginAndRegisterResponse{
			Response: Response{StatusCode: 200},
			UserId:   user.UserId,
			Token:    token,
		}
	}
}

func Login(username string, password string) UserLoginAndRegisterResponse {
	user := SelectUserByUsername(username)
	if user.UserId == 0 {
		return UserLoginAndRegisterResponse{
			Response: Response{
				StatusCode: 422,
				StatusMsg:  "用户不存在",
			},
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return UserLoginAndRegisterResponse{
			Response: Response{StatusCode: 1, StatusMsg: "用户名或密码错误"},
		}
	}
	token, err := CreateToken(user.UserId, user.Name)
	if err != nil {
		return UserLoginAndRegisterResponse{
			Response: Response{StatusCode: 500, StatusMsg: "服务器错误"},
		}
	}
	return UserLoginAndRegisterResponse{
		Response: Response{StatusCode: 0},
		UserId:   user.UserId,
		Token:    token,
	}
}
func GetUserInfo(userId int64) UserInfoResponse {
	user := SelectUserByUserId(userId)
	if user.UserId == 0 {
		return UserInfoResponse{
			Response: Response{
				StatusCode: 422,
				StatusMsg:  "用户不存在",
			},
		}
	}
	return UserInfoResponse{
		Response: Response{
			StatusCode: 200,
			StatusMsg:  "ok",
		},
		User: user,
	}
}
