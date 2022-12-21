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
