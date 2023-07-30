package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	router := gin.Default()

	//请求头的各种获取方式
	router.GET("/", func(c *gin.Context) {
		// 首字母大小写不区分  单词与单词之间用 - 连接
		// 用于获取一个请求头
		// 多个值返回切片
		fmt.Println(c.GetHeader("User-Agent"))
		fmt.Println(c.GetHeader("user-agent"))
		//fmt.Println(c.GetHeader("user-Agent"))
		//fmt.Println(c.GetHeader("user-AGent"))

		// Header 是一个普通的 map[string][]string
		fmt.Println(c.Request.Header)
		// 如果是使用 Get方法或者是 .GetHeader,那么可以不用区分大小写，并且返回第一个value
		fmt.Println(c.Request.Header.Get("User-Agent"))
		fmt.Println(c.Request.Header["User-Agent"])
		// 如果是用map的取值方式，请注意大小写问题
		fmt.Println(c.Request.Header["user-agent"])

		// 自定义的请求头，用Get方法也是免大小写
		fmt.Println(c.Request.Header.Get("Token"))
		fmt.Println(c.Request.Header.Get("token"))
		c.JSON(200, gin.H{"msg": "成功"})
	})

	//反爬虫的简易方法
	router.GET("/index", func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")
		//正则匹配
		if strings.Contains(userAgent, "python") {
			//爬虫来了
			c.JSON(0, gin.H{"data": "爬虫给爷爬，你个爬虫"})
			return
		}
		c.JSON(0, gin.H{"data": "爷的地盘禁止爬虫，你不是爬虫"})
	})

	// 设置响应头
	router.GET("/res", func(c *gin.Context) {
		c.Header("Token", "jhgeu%hsg845jUIF83jh")
		// 修改Content-Type的类型
		c.Header("Content-Type", "application/text; charset=utf-8")
		//c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(0, gin.H{"data": "看看响应头"})
	})

	router.Run(":8080")
}
