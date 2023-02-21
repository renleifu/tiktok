package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tiktok/main/config"
)

func main() {
	r := gin.Default()
	CollectRoutes(r)
	panic(r.Run(":" + strconv.Itoa(config.AppConfig.GetInt("server.port"))))
}
