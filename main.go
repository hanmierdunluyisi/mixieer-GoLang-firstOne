package main

import "github.com/gin-gonic/gin"

// 启动类
func main() {
	r := gin.Default() //携带基础中间件启动
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin!",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
