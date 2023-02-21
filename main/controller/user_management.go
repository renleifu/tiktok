package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
	"strconv"
	"tiktok/main/common"
	userService "tiktok/main/service"
)

type UserController struct {
}

type UserRegisterResponse struct {
	statusCode int32  `json:"status_code"`
	statusMsg  string `json:"status_msg"`
	userId     int64  `json:"user_id"`
	token      string `json:"token"`
}

func PostRegister(ctx *gin.Context) {
	//参数接收校验
	var username = ctx.Query("username")
	var password = ctx.Query("password")
	if username == "" || password == "" {
		ctx.JSON(http.StatusBadRequest, userService.UserLoginAndRegisterResponse{
			Response: common.Response{StatusCode: 422, StatusMsg: "用户名或密码不能为空"},
		})
		return
	}
	if !VerifyEmailFormat(username) {
		ctx.JSON(http.StatusBadRequest, userService.UserLoginAndRegisterResponse{
			Response: common.Response{StatusCode: 422, StatusMsg: "邮箱无效"},
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusBadRequest, userService.UserLoginAndRegisterResponse{
			Response: common.Response{StatusCode: 422, StatusMsg: "密码并不能小于6位"},
		})
		return
	}
	//password = utils.MD5WithSalt(password)
	encyptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, userService.UserLoginAndRegisterResponse{
			Response: common.Response{StatusCode: 500, StatusMsg: "服务器错误"},
		})
		return
	}
	var response = userService.Register(username, string(encyptedPassword))
	ctx.JSON(http.StatusOK, response)
}

func PostLogin(ctx *gin.Context) {
	var username = ctx.Query("username")
	var password = ctx.Query("password")
	if username == "" || password == "" {
		ctx.JSON(http.StatusBadRequest, userService.UserLoginAndRegisterResponse{
			Response: common.Response{StatusCode: 422, StatusMsg: "用户名或密码不能为空"},
		})
		return
	}
	if !VerifyEmailFormat(username) {
		ctx.JSON(http.StatusBadRequest, userService.UserLoginAndRegisterResponse{
			Response: common.Response{StatusCode: 422, StatusMsg: "邮箱无效"},
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusBadRequest, userService.UserLoginAndRegisterResponse{
			Response: common.Response{StatusCode: 422, StatusMsg: "密码并不能小于6位"},
		})
		return
	}

	response := userService.Login(username, password)
	ctx.JSON(http.StatusOK, response)
}
func GetUserInfo(ctx *gin.Context) {
	var userId, _ = strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	response := userService.GetUserInfo(userId)
	ctx.JSON(http.StatusOK, response)
}
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
