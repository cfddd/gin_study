package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	//增加了一个前缀“/api”
	//将一系列的路由放到一个组下，统一管理
	//例如，以下的路由前面统一加上api的前缀
	r := router.Group("/api")
	r.GET("/index", func(c *gin.Context) {
		c.String(200, "index")
	})
	r.GET("/home", func(c *gin.Context) {
		c.String(200, "home")
	})
	router.Run(":8080")
}
