package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
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

	r.Run()
}
