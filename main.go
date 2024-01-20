package main

import "github.com/gin-gonic/gin"

func Index(context *gin.Context) {
	context.String(200, "Hello world")

}

// web框架gin启动类
func main() {
	r := gin.Default() //携带基础中间件启动
	//
	r.GET("/test", Index)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080

}

// 连接mysql
