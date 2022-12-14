package routers

import (
	"gin-api-demo/pkg/settings"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())
	gin.SetMode(settings.GetConfig().RunMode)

	// TODO: 初始化日志

	// TODO: 初始化 MySQL

	// TODO: 初始化 Redis

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "PONG",
		})
	})

	return router
}
