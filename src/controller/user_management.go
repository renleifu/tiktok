package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tiktok/src/common"
	userService "tiktok/src/service"
	"tiktok/src/utils"
)

type UserController struct {
}

type UserRegisterResponse struct {
	statusCode int32  `json:"status_code"`
	statusMsg  string `json:"status_msg"`
	userId     int64  `json:"user_id"`
	token      string `json:"token"`
}

func (uc *UserController) PostRegister(ctx iris.Context) mvc.Result {
	//参数接收校验
	var username = ctx.URLParam("username")
	var password = ctx.URLParam("password")
	if username == "" || password == "" {
		return mvc.Response{
			Object: userService.UserLoginAndRegisterResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "用户名或密码不能为空"},
			},
		}
	}
	password = utils.MD5WithSalt(password)
	var response = userService.Register(username, password)

	return mvc.Response{
		Object: response,
	}
}

func (uc *UserController) PostLogin(ctx iris.Context) mvc.Response {
	var username = ctx.URLParam("username")
	var password = ctx.URLParam("password")
	if username == "" || password == "" {
		return mvc.Response{
			Object: userService.UserLoginAndRegisterResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "用户名或密码不能为空"},
			},
		}
	}
	response := userService.Login(username, utils.MD5WithSalt(password))

	return mvc.Response{
		Object: response,
	}
}

