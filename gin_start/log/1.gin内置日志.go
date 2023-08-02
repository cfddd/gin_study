package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	//启动gin，它会显示所有的路由，默认格式如下
	//gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	//	log.Printf(
	//		"[ CFD ] %v %v %v %v\n",
	//		httpMethod,
	//		absolutePath,
	//		handlerName,
	//		nuHandlers,
	//	)
	//}

	file, _ := os.Create("gin.log")

	//gin.DefaultWriter = io.MultiWriter(file)
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(os.Stdout, file)

	//改为release模式，可以不看debug日志
	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// router.Routes()它会返回已注册的路由列表,和上面自定义的一样
	fmt.Println(router.Routes())

	//匿名函数
	router.GET("/index", func(c *gin.Context) {})

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
