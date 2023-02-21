package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/main/controller"
	"tiktok/main/utils"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/douyin/user/register", controller.PostRegister)
	r.POST("/douyin/user/login", controller.PostLogin)

	r.Use(utils.JwtMiddleware())
	{
		r.POST("/douyin/user/", controller.GetUserInfo)
	}
	return r
}
