package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LogFormatterParams(c gin.LogFormatterParams) string {

	return fmt.Sprintf(
		"[CFD] %s | %s  %d  %s | %s  %s  %s | %s\n",
		c.TimeStamp.Format("2006-01-02-15:04:05"),
		c.StatusCodeColor(), c.StatusCode, c.ResetColor(),
		//设置颜色，内容，重置颜色
		c.MethodColor(), c.Method, c.ResetColor(),
		c.Path,
	)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	//两种调用自己的log显示
	//router.Use(gin.LoggerWithFormatter(LogFormatterParams))
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{Formatter: LogFormatterParams}))

	router.GET("/index", func(c *gin.Context) {})
	router.GET("/home", func(c *gin.Context) {})

	router.Run(":8080")
}
