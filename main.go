package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "12345")

		c.Next()

		latency := time.Since(t)
		log.Print(latency)

		status := c.Writer.Status()
		log.Println("-===", status)
	}
}

func Logger2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("name", "Lee")

		c.Next()

		log.Println("Logger2")
	}
}

func main() {

	r := gin.New()

	r.Use(Logger2(), Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		name := c.MustGet("name").(string)

		// 打印："12345"
		log.Println(example, "   ", name)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")

	// ==================================================  Demo1 ========================================
	//r := gin.Default()
	//r.LoadHTMLGlob("templates/**/*") // 多层目录下文件
	//r.GET("/", func(cxt *gin.Context) {
	//	cxt.JSON(http.StatusOK, gin.H{
	//		"message": "Hello Go",
	//	})
	//})
	//
	//r.GET("/someJSON", func(c *gin.Context) {
	//	data := map[string]interface{}{
	//		"lang": "Go 语言",
	//		"tag":  "<br/>",
	//	}
	//	c.AsciiJSON(http.StatusOK, data)
	//})
	//
	//r.GET("/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index/index.html", gin.H{
	//		"title": "Main website",
	//	})
	//})
	//
	//r.GET("/users/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
	//		"title": "Users",
	//	})
	//})
	//
	//r.Run()
	// ==================================================  Demo1 ========================================

}
