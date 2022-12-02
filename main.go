package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*") // 多层目录下文件
	r.GET("/", func(cxt *gin.Context) {
		cxt.JSON(http.StatusOK, gin.H{
			"message": "Hello Go",
		})
	})

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go 语言",
			"tag":  "<br/>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.html", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})

	r.Run()
}
